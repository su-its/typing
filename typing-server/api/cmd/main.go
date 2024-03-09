package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/api/presenter"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func main() {
	logger := slog.Default()

	// タイムゾーンの設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Error("failed to load location: %v", err)
		return
	}

	var addr = os.Getenv("DB_ADDR")
	if addr == "" {
		addr = "db:3306" // アドレス（Docker Compose内でのサービス名とポート）
	}

	// MySQLの接続設定
	mysqlConfig := &mysql.Config{
		DBName:    "typing-db", // データベース名
		User:      "user",      // ユーザー名
		Passwd:    "password",  // パスワード
		Net:       "tcp",       // ネットワークタイプ
		Addr:      addr,
		ParseTime: true, // 時刻をtime.Timeで解析する
		Loc:       jst,  // タイムゾーン
	}

	// entクライアントの初期化
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		logger.Error("failed to open ent client: %v", err)
		return
	}
	defer entClient.Close()
	logger.Info("ent client is opened")

	// スキーマの作成
	if err := entClient.Schema.Create(context.Background()); err != nil {
		logger.Error("failed to create schema: %v", err)
		return
	}
	logger.Info("schema is created")

	// ルートの登録
	presenter.RegisterRoutes()

	// WaitGroupの宣言
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
		// サーバーの設定
		server := &http.Server{Addr: ":8080"}
		// 非同期でサーバーを開始
		go func() {
			logger.Info("server is running at Addr :8080")
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				logger.Error("failed to listen and serve: %v", err)
				errChan <- err // エラーをチャネルに送信
			}
		}()
		// シグナルを待機
		<-sigChan
		logger.Info("shutting down the server...")
		ctx := context.TODO() // Use context.TODO() as a temporary placeholder
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("error during server shutdown: %v", err)
			errChan <- err // エラーをチャネルに送信
		}
	}()
	select {
	case <-errChan: // エラーが発生した場合
		logger.Error("server stopped due to an error")
	case sig := <-sigChan: // シグナルを受信した場合
		logger.Info("received signal: %s", sig)
	}
	wg.Wait() // HTTPサーバーの終了を待機
	close(errChan)
	logger.Info("server exited")
}
