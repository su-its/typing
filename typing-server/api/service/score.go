package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/repository"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
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

func GetMyScoreRanking(ctx context.Context, client *ent.Client, userID uuid.UUID, sortBy string) (int, error) {
	// ユーザーの最大スコアを取得
	userMaxScore, err := repository.GetMaxScoreByUserID(ctx, client, userID, sortBy)
	if err != nil {
		if ent.IsNotFound(err) {
			// ユーザーのスコアが存在しない場合は、ランキング外として0を返す
			return 0, nil
		}
		return 0, err
	}

	// ユーザーの最大スコアより上位のスコアをカウント
	rank, err := repository.CountHigherScores(ctx, client, userMaxScore, sortBy)
	if err != nil {
		return 0, err
	}

	// ランキングを返す
	return rank + 1, nil
}
