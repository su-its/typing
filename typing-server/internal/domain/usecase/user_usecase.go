package usecase

import (
	"context"
	"errors"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

type IUserUseCase interface {
	GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error)
	CreateUser(ctx context.Context, studentNumber string, handleName string) (*model.User, error)
}

// コンパイル時にインターフェースの実装を確認
var _ IUserUseCase = (*UserUseCase)(nil)

// UserUseCase はユーザー関連のユースケースを管理する
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase は UserUseCase のインスタンスを生成する
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// GetUserByStudentNumber は学籍番号を元にユーザーを取得するユースケース
func (uc *UserUseCase) GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error) {
	return uc.userRepo.GetUserByStudentNumber(ctx, studentNumber)
}

// CreateUser はユーザーを作成するユースケース
// ユーザーが既に存在する場合はErrUserAlreadyExistsを返す
func (uc *UserUseCase) CreateUser(ctx context.Context, studentNumber string, handleName string) (*model.User, error) {
	u, err := uc.userRepo.CreateUser(ctx, studentNumber, handleName)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrUserAlreadyExists):
			return nil, ErrUserAlreadyExists
		default:
			return nil, err
		}
	}
	return u, nil
}
