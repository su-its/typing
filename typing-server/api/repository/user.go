package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/user"
)

func GetUserByStudentNumber(ctx context.Context, studentNumber string) (*ent.User, error) {
	client := ent.FromContext(ctx)

	entUser, err := client.User.Query().
		Where(user.StudentNumber(studentNumber)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &ent.User{
		ID:            entUser.ID,
		StudentNumber: entUser.StudentNumber,
		HandleName:    entUser.HandleName,
	}, nil
}
