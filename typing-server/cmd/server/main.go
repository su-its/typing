package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/internal/config"
	"github.com/su-its/typing/typing-server/internal/router"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/pkg/logger"
)

func main() {
	// ロガーの初期化
	log := logger.New()
	config := config.New(log)

	// タイムゾーンの設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Error("failed to load timezone",
			"error", err,
			"timezone", "Asia/Tokyo")
		return
	}

	// MySQLの接続設定
	mysqlConfig := &mysql.Config{
		DBName:    "typing-db",
		User:      "user",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      config.DBAddr,
		ParseTime: true,
		Loc:       jst,
	}

	// entクライアントの初期化
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
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

	// WaitGroupとチャネルの初期化
	var wg sync.WaitGroup
	// エラーを通知するためのチャネル
	errChan := make(chan error, 1)
	// シグナルハンドリングの準備
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// HTTPサーバーの非同期起動
	wg.Add(1)
	go func() {
		defer wg.Done() // 関数終了時にWaitGroupをデクリメント

		// ルーティングの設定
		r := router.SetupRouter(log, entClient, config)

		// サーバーの設定
		server := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}

		// 非同期でサーバーを起動
		go func() {
			log.Info("starting HTTP server",
				"addr", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Error("server failed to start",
					"error", err)
				errChan <- err // エラーをチャネルに送信
			}
		}()

		// エラーまたはシグナルを待機
		select {
		case err := <-errChan:
			log.Error("server stopped due to error",
				"error", err)
		case sig := <-sigChan:
			log.Info("received shutdown signal",
				"signal", sig)

			// グレースフルシャットダウン
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				log.Error("error during server shutdown",
					"error", err)
				errChan <- err // エラーをチャネルに送信
			}
		}
	}()

	wg.Wait() // HTTPサーバーの終了を待機
	close(errChan)
	log.Info("server shutdown completed")
}
