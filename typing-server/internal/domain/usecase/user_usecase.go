package usecase

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

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
