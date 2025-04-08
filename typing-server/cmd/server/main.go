package main

import (
	"context"
	"net/http"

	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/domain/service"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/repository"
	"github.com/su-its/typing/typing-server/internal/interfaces"
	"github.com/su-its/typing/typing-server/internal/interfaces/handler"
	"github.com/su-its/typing/typing-server/pkg/logger"
	"github.com/su-its/typing/typing-server/pkg/middleware"
)

func main() {
	log := logger.New()
	cfg, err := config.New()
	if err != nil {
		log.Error("failed to load config", "error", err)
		return
	}

	log.Info("config",
		"environment", cfg.Environment,
		"db_addr", cfg.DBAddr)
	log.Info("timezone",
		"timezone", cfg.GetLocation().String())

	log.Info("mysql config", "dsn", cfg.GetMySQLDSN())

	// entクライアントの初期化
	entClient, err := ent_generated.Open("mysql", cfg.GetMySQLDSN())
	if err != nil {
		log.Error("failed to open database connection",
			"error", err,
			"dsn", cfg.GetMySQLDSN())
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
	userHandler := handler.NewUserHandler(userUseCase, log)
	scoreHandler := handler.NewScoreHandler(scoreUseCase, log)

	// ルーターの作成
	router := interfaces.NewRouter(healthHandler, userHandler, scoreHandler, cfg)

	// サーバー起動
	log.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", middleware.LoggingMiddleware(router)); err != nil {
		log.Error("failed to start server",
			"error", err)
	}
}
