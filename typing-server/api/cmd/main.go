package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/api/presenter"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func main() {
	logger := slog.Default()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Error("failed to load location", fmt.Errorf("error: %w", err))
	}

	mysqlConfig := &mysql.Config{
		DBName:    "typing-db",
		User:      "user",
		Passwd:    "password",       // 環境変数から取得するか、直接指定
		Addr:      "db:33061", // Docker Compose内でのサービス名とポート
		ParseTime: true,
		Loc:       jst,
	}
	
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		logger.Error("failed to open ent client", fmt.Errorf("error: %w", err))
	} else {
		logger.Info("ent client is opened")
	}

	if err := entClient.Schema.Create(context.Background()); err != nil {
		logger.Error("failed to create schema", fmt.Errorf("error: %w", err))
	} else {
		logger.Info("schema is created")
	}

	presenter.RegisterRoutes()
	go func() {
		logger.Info("server is running")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			logger.Error("failed to listen and serve", fmt.Errorf("error: %w", err))
		}
	}()
}
