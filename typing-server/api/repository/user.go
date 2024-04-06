package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/score"
	"github.com/su-its/typing/typing-server/domain/repository/ent/user"
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

func GetMaxScoreByUserID(ctx context.Context, client *ent.Client, userID uuid.UUID, sortBy string) (*ent.Score, error) {
	var maxScore *ent.Score
	var err error

	switch sortBy {
	case "accuracy":
		maxScore, err = client.Score.Query().
			Where(
				score.And(
					score.UserID(userID),
					score.KeystrokesGTE(120),
					score.AccuracyGTE(0.95),
				),
			).
			Order(ent.Desc(score.FieldAccuracy)).
			First(ctx)
	case "keystrokes":
		maxScore, err = client.Score.Query().
			Where(
				score.And(
					score.UserID(userID),
					score.KeystrokesGTE(120),
					score.AccuracyGTE(0.95),
				),
			).
			Order(ent.Desc(score.FieldKeystrokes)).
			First(ctx)
	default:
		return nil, fmt.Errorf("invalid sort by parameter: %s", sortBy)
	}

	if err != nil {
		return nil, err
	}

	return maxScore, nil
}
