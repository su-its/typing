package userimport

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

type stubUserRepository struct {
	createFn func(ctx context.Context, studentNumber string, handleName string) (*model.User, error)
}

func (s stubUserRepository) GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error) {
	return nil, nil
}

func (s stubUserRepository) CreateUser(ctx context.Context, studentNumber string, handleName string) (*model.User, error) {
	return s.createFn(ctx, studentNumber, handleName)
}

func TestSeedUsers(t *testing.T) {
	t.Parallel()

	t.Run("新規作成と既存スキップを集計できる", func(t *testing.T) {
		t.Parallel()

		callCount := 0
		repo := stubUserRepository{
			createFn: func(ctx context.Context, studentNumber string, handleName string) (*model.User, error) {
				callCount++
				if studentNumber == "726A0002" {
					return nil, repository.ErrUserAlreadyExists
				}
				return &model.User{StudentNumber: studentNumber, HandleName: handleName}, nil
			},
		}

		summary, err := SeedUsers(context.Background(), repo, []model.User{
			{StudentNumber: "726A0001", HandleName: "青山　哲"},
			{StudentNumber: "726A0002", HandleName: "浅井　秀羽"},
		})
		if err != nil {
			t.Fatalf("SeedUsers() error = %v", err)
		}
		if callCount != 2 {
			t.Fatalf("SeedUsers() callCount = %d, want 2", callCount)
		}
		if summary.Total != 2 || summary.Created != 1 || summary.Skipped != 1 {
			t.Fatalf("SeedUsers() summary = %+v, want Total=2 Created=1 Skipped=1", summary)
		}
	})

	t.Run("予期しないエラーは返す", func(t *testing.T) {
		t.Parallel()

		repo := stubUserRepository{
			createFn: func(ctx context.Context, studentNumber string, handleName string) (*model.User, error) {
				return nil, errors.New("db is down")
			},
		}

		_, err := SeedUsers(context.Background(), repo, []model.User{
			{StudentNumber: "726A0001", HandleName: "青山　哲"},
		})
		if err == nil {
			t.Fatal("SeedUsers() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "create user") {
			t.Fatalf("SeedUsers() error = %q, want substring %q", err.Error(), "create user")
		}
	})
}
