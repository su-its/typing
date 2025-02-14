package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/app/repository"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/infra/ent"
)

func GetScoresRanking(ctx context.Context, client *ent.Client, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
	response, err := repository.GetScoresRanking(ctx, client, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func CreateScore(ctx context.Context, client *ent.Client, userID uuid.UUID, keystrokes int, accuracy float64) error {
	if err := repository.CreateScore(ctx, client, userID, keystrokes, accuracy); err != nil {
		return err
	}

	return nil
}
