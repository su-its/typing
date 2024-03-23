package service

import (
	"context"

	"github.com/su-its/typing/typing-server/api/repository"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func GetUserByStudentNumber(ctx context.Context, studentNumber string) (*ent.User, error) {
	user, err := repository.GetUserByStudentNumber(ctx, studentNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
}
