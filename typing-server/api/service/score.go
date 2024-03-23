package service

import (
	"context"

	"github.com/su-its/typing/typing-server/api/repository"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func GetScoresRanking(ctx context.Context, sortBy string, start, limit int) ([]*model.ScoreRanking, error) {
	rankings, err := repository.GetScoresRanking(ctx, sortBy, start, limit)
	if err != nil {
		return nil, err
	}

	return rankings, nil
}

func CreateScore(ctx context.Context, score *ent.Score) error {
	if err := repository.CreateScore(ctx, score); err != nil {
		return err
	}

	return nil
}
