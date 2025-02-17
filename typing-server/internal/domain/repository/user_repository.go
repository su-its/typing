package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/domain/model"
)

type UserRepository interface {
	// GetUserByStudentNumber は、指定された学籍番号を持つユーザーを取得する。
	// 該当するユーザーが存在しない場合は、(nil, nil) を返す。
	GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error)
	// CreateUser は、指定された学籍番号とハンドルネームを持つユーザーを作成する。
	// 既に同じ学籍番号を持つユーザーが存在する場合は、{nil, ErrAlreadyExists} を返す。
	CreateUser(ctx context.Context, studentNumber string, handleName string) (*model.User, error)
}
