package userimport

import (
	"context"
	"errors"
	"fmt"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

type Summary struct {
	Total   int
	Created int
	Skipped int
}

func SeedUsers(ctx context.Context, userRepo repository.UserRepository, users []model.User) (Summary, error) {
	summary := Summary{Total: len(users)}

	for _, user := range users {
		if user.StudentNumber == "" {
			return summary, errors.New("student number is empty")
		}
		if user.HandleName == "" {
			return summary, fmt.Errorf("handle name is empty for student number %q", user.StudentNumber)
		}

		_, err := userRepo.CreateUser(ctx, user.StudentNumber, user.HandleName)
		switch {
		case err == nil:
			summary.Created++
		case errors.Is(err, repository.ErrUserAlreadyExists):
			summary.Skipped++
		default:
			return summary, fmt.Errorf("create user %q: %w", user.StudentNumber, err)
		}
	}

	return summary, nil
}
