package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	entrepo "github.com/su-its/typing/typing-server/internal/infra/ent/repository"
	"github.com/su-its/typing/typing-server/internal/seed/userimport"
	"github.com/su-its/typing/typing-server/pkg/logger"
)

func main() {
	var csvPath string
	var createSchema bool

	flag.StringVar(&csvPath, "csv", "", "インポートするCSVファイルのパス")
	flag.BoolVar(&createSchema, "createSchema", true, "起動時にスキーマを作成する")
	flag.Parse()

	if csvPath == "" {
		fmt.Println("エラー: -csv は必須です")
		flag.Usage()
		os.Exit(1)
	}

	logr := logger.New()
	cfg := config.New()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logr.Error("タイムゾーンのロードに失敗", "error", err, "timezone", "Asia/Tokyo")
		os.Exit(1)
	}

	users, err := userimport.LoadUsersFromCSV(csvPath)
	if err != nil {
		logr.Error("CSVの読み込みに失敗", "error", err, "csv", csvPath)
		os.Exit(1)
	}

	mysqlConfig := &mysql.Config{
		DBName:               "typing-db",
		User:                 "user",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 cfg.DBAddr,
		ParseTime:            true,
		Loc:                  jst,
		AllowNativePasswords: true,
		TLSConfig:            "false",
	}

	entClient, err := ent_generated.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		logr.Error("データベース接続に失敗", "error", err, "config", mysqlConfig.FormatDSN())
		os.Exit(1)
	}
	defer entClient.Close()

	ctx := context.Background()
	if createSchema {
		if err := entClient.Schema.Create(ctx); err != nil {
			logr.Error("スキーマ作成に失敗", "error", err)
			os.Exit(1)
		}
	}

	userRepo := entrepo.NewEntUserRepository(entClient)
	summary, err := userimport.SeedUsers(ctx, userRepo, users)
	if err != nil {
		logr.Error("ユーザーインポートに失敗", "error", err, "csv", csvPath)
		os.Exit(1)
	}

	logr.Info(
		"ユーザーインポート完了",
		"environment", cfg.Environment,
		"db_addr", cfg.DBAddr,
		"csv", csvPath,
		"total", summary.Total,
		"created", summary.Created,
		"skipped", summary.Skipped,
	)
}
