package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/su-its/typing/typing-server/api/handler"
	"github.com/su-its/typing/typing-server/api/router"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/user"
)

func main() {
	seedFlag := flag.Bool("seed", false, "シードデータを挿入する場合はtrueを指定")
	flag.Parse()

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
	handler.SetEntClient(entClient)
	logger.Info("ent client is opened")

	// スキーマの作成
	if err := entClient.Schema.Create(context.Background()); err != nil {
		logger.Error("failed to create schema: %v", err)
		return
	}
	logger.Info("schema is created")

	// シードデータの挿入
	if *seedFlag {
		if err := seedData(context.Background(), entClient, logger); err != nil {
			logger.Error("failed to seed data: %v", err)
			return
		}
		logger.Info("シードデータが挿入されました")
	} else {
		logger.Info("シードデータは挿入されませんでした")
	}

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
		// ルーティングの設定
		r := router.SetupRouter()

		// サーバーの設定
		server := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
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

func seedData(ctx context.Context, client *ent.Client, logger *slog.Logger) error {
	// シードデータの作成
	for i := 0; i < 10; i++ {
		studentNumber := fmt.Sprintf("user%d", i+1)
		handleName := fmt.Sprintf("handle%d", i+1)

		isAlreadySeeded, err := client.User.Query().Where(user.StudentNumber(studentNumber)).Exist(ctx)
		if err != nil {
			return err
		}
		if isAlreadySeeded {
			logger.Info("User with student number already seeded, skipping", slog.String("studentNumber", studentNumber))
			continue
		}

		u, err := client.User.Create().
			SetStudentNumber(studentNumber).
			SetHandleName(handleName).
			Save(ctx)
		if err != nil {
			panic(err)
		}

		logger.Info("Created user", slog.String("studentNumber", studentNumber), slog.String("handleName", handleName))

		var maxKeystrokesScore, maxAccuracyScore *ent.Score

		for j := 0; j < 5; j++ {
			keystrokes := rand.Intn(200) + 100
			accuracy := rand.Float64()

			s, err := client.Score.Create().
				SetKeystrokes(keystrokes).
				SetAccuracy(accuracy).
				SetCreatedAt(time.Now()).
				SetUser(u).
				Save(ctx)
			if err != nil {
				panic(err)
			}

			logger.Info("Created score", slog.Int("keystrokes", keystrokes), slog.Float64("accuracy", accuracy), slog.String("studentNumber", studentNumber))

			if s.Keystrokes < 120 || s.Accuracy < 0.95 {
				continue
			}

			if maxKeystrokesScore == nil || s.Keystrokes > maxKeystrokesScore.Keystrokes {
				maxKeystrokesScore = s
			}
			if maxAccuracyScore == nil || s.Accuracy > maxAccuracyScore.Accuracy {
				maxAccuracyScore = s
			}
		}

		// 最大のKeystrokesスコアと最大のAccuracyスコアのフラグを設定
		if maxKeystrokesScore != nil {
			err = maxKeystrokesScore.Update().SetIsMaxKeystrokes(true).Exec(ctx)
			if err != nil {
				return err
			}
			logger.Info("Set is_max_keystrokes flag", slog.Int("keystrokes", maxKeystrokesScore.Keystrokes), slog.String("studentNumber", studentNumber))
		}
		if maxAccuracyScore != nil {
			err = maxAccuracyScore.Update().SetIsMaxAccuracy(true).Exec(ctx)
			if err != nil {
				return err
			}
			logger.Info("Set is_max_accuracy flag", slog.Float64("accuracy", maxAccuracyScore.Accuracy), slog.String("studentNumber", studentNumber))
		}
	}

	return nil
}
