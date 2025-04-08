package main

import (
	"context"
	"encoding/csv"
	"errors"
	"log/slog"
	"os"

	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/repository"
	"github.com/su-its/typing/typing-server/pkg/logger"
)

func main() {
	log := logger.New()

	cfg, err := loadConfig(log)
	if err != nil {
		return
	}

	entClient, err := connectDatabase(cfg, log)
	if err != nil {
		return
	}
	defer entClient.Close()

	userRepo := repository.NewEntUserRepository(entClient)
	userUseCase := usecase.NewUserUseCase(userRepo)

	users, err := parseUsersFromCSV("users.csv")
	if err != nil {
		log.Error("CSV読み込み失敗", "error", err)
		return
	}
	log.Info("CSV読み込み成功", "count", len(users))

	if err := createUsers(context.Background(), userUseCase, users, log); err != nil {
		log.Error("ユーザー作成失敗", "error", err)
		return
	}

	log.Info("ユーザー作成完了", "count", len(users))
}

func loadConfig(log *slog.Logger) (*config.Config, error) {
	cfg, err := config.New()
	if err != nil {
		log.Error("設定ファイル読み込み失敗", "error", err)
		return nil, err
	}
	log.Info("設定ファイル読み込み成功", "environment", cfg.Environment)
	return cfg, nil
}

func connectDatabase(cfg *config.Config, log *slog.Logger) (*ent_generated.Client, error) {
	client, err := ent_generated.Open("mysql", cfg.GetMySQLDSN())
	if err != nil {
		log.Error("DB接続失敗", "error", err, "dsn", cfg.GetMySQLDSN())
		return nil, err
	}
	log.Info("DB接続成功")
	return client, nil
}

type csvUser struct {
	StudentNumber string
	HandleName    string
}

func parseUsersFromCSV(path string) ([]*csvUser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	users := make([]*csvUser, 0, len(records))
	for _, record := range records {
		if len(record) < 2 {
			continue // 空行や不完全なレコードをスキップ
		}
		users = append(users, &csvUser{
			StudentNumber: record[0],
			HandleName:    record[1],
		})
	}
	return users, nil
}

func createUsers(ctx context.Context, useCase *usecase.UserUseCase, users []*csvUser, log *slog.Logger) error {
	for _, user := range users {
		if _, err := useCase.CreateUser(ctx, user.StudentNumber, user.HandleName); err != nil {
			log.Error("単一ユーザー作成失敗", "student_number", user.StudentNumber, "error", err)
			if errors.Is(err, usecase.ErrUserAlreadyExists) {
				log.Warn("ユーザーが既に存在します", "student_number", user.StudentNumber)
			}
			return err
		}
	}
	return nil
}
