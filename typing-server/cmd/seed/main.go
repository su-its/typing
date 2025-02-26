package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/domain/service"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/repository"
	"github.com/su-its/typing/typing-server/pkg/logger"
)

func main() {
	// コマンドライン引数の定義
	var numUsers int
	var numScores int
	var minKeystrokes int
	var minAccuracy float64
	var maxConcurrent int
	flag.IntVar(&numUsers, "users", 0, "シードするユーザー数")
	flag.IntVar(&numScores, "scores", 0, "1ユーザーあたりにシードするスコア数")
	flag.IntVar(&minKeystrokes, "minKeystrokes", 0, "各ユーザーに登録する最低スコア（keystrokes）の下限値")
	flag.Float64Var(&minAccuracy, "minAccuracy", 0.95, "各ユーザーに登録するaccuracyの最低値（0.0〜1.0）")
	flag.IntVar(&maxConcurrent, "concurrent", 10, "同時実行するgoroutineの最大数")
	flag.Parse()

	if numUsers <= 0 || numScores <= 0 || minKeystrokes <= 0 {
		fmt.Println("エラー: -users と -scores と -minKeystrokes は正の整数で指定してください")
		flag.Usage()
		os.Exit(1)
	}
	if minAccuracy < 0.0 || minAccuracy >= 1.0 {
		fmt.Println("エラー: -minAccuracy は0.0以上1.0未満で指定してください")
		flag.Usage()
		os.Exit(1)
	}

	// ログ・設定の初期化
	logr := logger.New()
	cfg := config.New()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logr.Error("タイムゾーンのロードに失敗", "error", err, "timezone", "Asia/Tokyo")
		os.Exit(1)
	}
	logr.Info("config", "environment", cfg.Environment, "db_addr", cfg.DBAddr)
	logr.Info("timezone", "timezone", "Asia/Tokyo")

	// MySQL設定
	mysqlConfig := &mysql.Config{
		DBName:               "typing-db",
		User:                 "user",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 cfg.DBAddr,
		ParseTime:            true,
		Loc:                  jst,
		AllowNativePasswords: true,
		TLSConfig:            "false", // SSL/TLS接続を無効化
	}
	logr.Info("mysql config", "config", mysqlConfig.FormatDSN())

	// entクライアントの初期化
	entClient, err := ent_generated.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		logr.Error("データベース接続に失敗", "error", err, "config", mysqlConfig.FormatDSN())
		os.Exit(1)
	}
	defer entClient.Close()
	logr.Info("DB接続確立")

	// スキーマ作成
	ctx := context.Background()
	if err := entClient.Schema.Create(ctx); err != nil {
		logr.Error("スキーマ作成に失敗", "error", err)
		os.Exit(1)
	}
	logr.Info("スキーマ作成成功")

	// 各種リポジトリ・ユースケースの初期化
	txManager := repository.NewEntTxManager(entClient)
	userRepo := repository.NewEntUserRepository(entClient)
	scoreRepo := repository.NewEntScoreRepository(entClient)
	scoreService := service.NewScoreService(scoreRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	scoreUseCase := usecase.NewScoreUseCase(txManager, scoreRepo, scoreService)

	var wg sync.WaitGroup

	// ユーザー作成とスコア登録処理を並列化
	// 同時実行数を制限するセマフォを追加
	sem := make(chan struct{}, maxConcurrent)

	logr.Info("シード作業開始", "numUsers", numUsers, "numScores", numScores, "minKeystrokes", minKeystrokes, "minAccuracy", minAccuracy)

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go func(i int) {
			// セマフォを取得
			sem <- struct{}{}
			defer func() {
				// 処理完了時にセマフォを解放
				<-sem
				wg.Done()
			}()

			// 8桁のゼロパディングされた学生番号を生成（00000000から開始）
			studentNumber := fmt.Sprintf("%08d", i-1)
			handleName := fmt.Sprintf("User%d", i)

			// ユーザー作成
			user, err := userUseCase.CreateUser(ctx, studentNumber, handleName)
			if err != nil {
				// ユーザーが既に存在する場合はログに出力して処理継続
				if errors.Is(err, usecase.ErrUserAlreadyExists) {
					// ユーザーが既に存在する場合は処理をスキップ
					return
				}
				logr.Error("ユーザー作成に失敗", "studentNumber", studentNumber, "error", err)
				return
			}

			uid, err := uuid.Parse(user.ID)
			if err != nil {
				logr.Error("ユーザーIDのパースに失敗", "user.ID", user.ID, "error", err)
				return
			}

			// ユーザーごとにnumScores分のスコアを生成
			var scoreWg sync.WaitGroup
			for j := 0; j < numScores; j++ {
				scoreWg.Add(1)
				go func() {
					defer scoreWg.Done()

					// ランダムなスコア生成
					// keystrokesはminKeystrokes以上、minKeystrokesからminKeystrokes+200の範囲
					keystrokes := minKeystrokes + rand.Intn(200)
					// accuracyはminAccuracy〜1.0の範囲
					accuracy := minAccuracy + rand.Float64()*(1.0-minAccuracy)

					// スコア登録
					err := scoreUseCase.RegisterScore(ctx, uid, keystrokes, accuracy)
					if err != nil {
						logr.Error("スコア登録に失敗", "user.ID", user.ID, "error", err)
					}
				}()
			}

			// 全てのスコア登録が完了するのを待つ
			scoreWg.Wait()
		}(i)
	}

	wg.Wait()
	logr.Info("シード処理完了")
}
