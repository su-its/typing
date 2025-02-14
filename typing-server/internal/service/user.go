package service

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/repository"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func GetUserByStudentNumber(ctx context.Context, client *ent.Client, studentNumber string) (*ent.User, error) {
	user, err := repository.GetUserByStudentNumber(ctx, client, studentNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
}
