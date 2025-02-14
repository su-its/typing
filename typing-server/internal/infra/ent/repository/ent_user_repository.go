package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
	"github.com/su-its/typing/typing-server/internal/infra/ent/generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/generated/user"
)

// EntUserRepository は ent を使用して UserRepository を実装する。
type EntUserRepository struct {
	client *generated.Client
}

// コンパイル時チェック: EntUserRepository が UserRepository インターフェースを実装していることを保証する。
var _ repository.UserRepository = (*EntUserRepository)(nil)

// NewEntUserRepository は EntUserRepository のコンストラクタ。
func NewEntUserRepository(client *generated.Client) *EntUserRepository {
	return &EntUserRepository{client: client}
}

// GetUserByStudentNumber は、指定された学籍番号を持つユーザーを取得する。
// 該当するユーザーが存在しない場合は、(nil, nil) を返す。
func (r *EntUserRepository) GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error) {
	entUser, err := r.client.User.Query().
		WithScores().
		Where(user.StudentNumberEQ(studentNumber)).
		Only(ctx)

	if generated.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &model.User{
		ID:            entUser.ID.String(),
		StudentNumber: entUser.StudentNumber,
		HandleName:    entUser.HandleName,
		CreatedAt:     entUser.CreatedAt,
		UpdatedAt:     entUser.UpdatedAt,
	}, nil
}
