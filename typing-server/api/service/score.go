package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/repository"
	"github.com/su-its/typing/typing-server/domain/model"
)

func GetScoresRanking(ctx context.Context, sortBy string, start int) ([]*model.ScoreRanking, error) {
	rankings, err := repository.GetScoresRanking(ctx, sortBy, start)
	if err != nil {
		return nil, err
	}

	return rankings, nil
}

func CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
	if err := repository.CreateScore(ctx, userID, keystrokes, accuracy); err != nil {
		return err
	}

	return nil
}
