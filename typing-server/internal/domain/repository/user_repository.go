package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/domain/model"
)

type UserRepository interface {
	// GetUserByStudentNumber は、指定された学籍番号を持つユーザーを取得する。
	// 該当するユーザーが存在しない場合は、(nil, nil) を返す。
	GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error)
}
