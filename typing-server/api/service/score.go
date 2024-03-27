package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/repository"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func GetScoresRanking(ctx context.Context, client *ent.Client, sortBy string, start, limit int) ([]*model.ScoreRanking, error) {
	rankings, err := repository.GetScoresRanking(ctx, client, sortBy, start, limit)
	if err != nil {
		return nil, err
	}

	return rankings, nil
}

func CreateScore(ctx context.Context, client *ent.Client, userID uuid.UUID, keystrokes int, accuracy float64) error {
	if err := repository.CreateScore(ctx, client, userID, keystrokes, accuracy); err != nil {
		return err
	}

	return nil
}
