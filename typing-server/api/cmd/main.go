package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/api/presenter"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func main() {
	// 標準のログパッケージを使用
	logger := log.Default()

	// タイムゾーンの設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Printf("failed to load location: %v", err)
		return
	}

	// MySQLの接続設定
	mysqlConfig := &mysql.Config{
		DBName:    "typing-db", // データベース名
		User:      "user",      // ユーザー名
		Passwd:    "password",  // パスワード
		Net:       "tcp",       // ネットワークタイプ
		Addr:      "db:3306",   // アドレス（Docker Compose内でのサービス名とポート）
		ParseTime: true,        // 時刻をtime.Timeで解析する
		Loc:       jst,         // タイムゾーン
	}

	// entクライアントの初期化
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		logger.Printf("failed to open ent client: %v", err)
		return
	} else {
		logger.Println("ent client is opened")
	}

	// スキーマの作成
	if err := entClient.Schema.Create(context.Background()); err != nil {
		logger.Printf("failed to create schema: %v", err)
		return
	} else {
		logger.Println("schema is created")
	}

	// ルートの登録
	presenter.RegisterRoutes()

	// WaitGroupの宣言
	var wg sync.WaitGroup

	// HTTPサーバーの非同期起動
	wg.Add(1)
	go func() {
		defer wg.Done() // 関数終了時にWaitGroupをデクリメント

		logger.Println("server is running at http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			logger.Printf("failed to listen and serve: %v", err)
		}
	}()

	wg.Wait() // HTTPサーバーの終了を待機
}
