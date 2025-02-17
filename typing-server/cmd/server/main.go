package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/domain/service"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/repository"
	"github.com/su-its/typing/typing-server/internal/interfaces"
	"github.com/su-its/typing/typing-server/internal/interfaces/handler"
	"github.com/su-its/typing/typing-server/pkg/logger"
)

func main() {
	log := logger.New()
	config := config.New()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Error("failed to load timezone",
			"error", err,
			"timezone", "Asia/Tokyo")
		return
	}
	log.Info("config",
		"environment", config.Environment,
		"db_addr", config.DBAddr)
	log.Info("timezone",
		"timezone", "Asia/Tokyo")

	mysqlConfig := &mysql.Config{
		DBName:    "typing-db",
		User:      "user",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      config.DBAddr,
		ParseTime: true,
		Loc:       jst,
	}
	log.Info("mysql config",
		"config", mysqlConfig.FormatDSN())

	// entクライアントの初期化
	entClient, err := ent_generated.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Error("failed to open database connection",
			"error", err,
			"config", mysqlConfig.FormatDSN())
		return
	}
	defer entClient.Close()
	log.Info("database connection established")

	// スキーマの作成
	ctx := context.Background()
	if err := entClient.Schema.Create(ctx); err != nil {
		log.Error("failed to create database schema",
			"error", err)
		return
	}
	log.Info("database schema created successfully")

	// トランザクションマネージャの作成
	txManager := repository.NewEntTxManager(entClient)

	// リポジトリの作成
	userRepo := repository.NewEntUserRepository(entClient)
	scoreRepo := repository.NewEntScoreRepository(entClient)

	// サービスの作成
	scoreService := service.NewScoreService(scoreRepo)

	// ユースケースの作成
	userUseCase := usecase.NewUserUseCase(userRepo)
	scoreUseCase := usecase.NewScoreUseCase(txManager, scoreRepo, scoreService)

	// ハンドラの作成
	healthHandler := handler.NewHealthCheckHandler()
	userHandler := handler.NewUserHandler(userUseCase)
	scoreHandler := handler.NewScoreHandler(scoreUseCase)

	// ルーターの作成
	router := interfaces.NewRouter(healthHandler, userHandler, scoreHandler, config)

	// サーバー起動
	log.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Error("failed to start server",
			"error", err)
	}
}
