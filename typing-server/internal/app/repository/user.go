package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/infra/ent"
	"github.com/su-its/typing/typing-server/internal/infra/ent/user"
)

func GetUserByStudentNumber(ctx context.Context, client *ent.Client, studentNumber string) (*ent.User, error) {
	entUser, err := client.User.Query().
		WithScores().
		Where(user.StudentNumberEQ(studentNumber)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &ent.User{
		ID:            entUser.ID,
		StudentNumber: entUser.StudentNumber,
		HandleName:    entUser.HandleName,
		CreatedAt:     entUser.CreatedAt,
		UpdatedAt:     entUser.UpdatedAt,
		Edges:         entUser.Edges,
	}, nil
}
