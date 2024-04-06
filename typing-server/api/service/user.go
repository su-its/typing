package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/repository"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/score"
)

func GetUserByStudentNumber(ctx context.Context, client *ent.Client, studentNumber string) (*ent.User, error) {
	user, err := repository.GetUserByStudentNumber(ctx, client, studentNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
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
	var rank int

	switch sortBy {
	case "accuracy":
		rank, err = client.Score.Query().
			Where(
				score.And(
					score.KeystrokesGTE(120),
					score.AccuracyGTE(0.95),
					score.Or(
						score.AccuracyGT(userMaxScore.Accuracy),
						score.And(
							score.AccuracyEQ(userMaxScore.Accuracy),
							score.KeystrokesGT(userMaxScore.Keystrokes),
						),
					),
				),
			).
			Count(ctx)
	case "keystrokes":
		rank, err = client.Score.Query().
			Where(
				score.And(
					score.KeystrokesGTE(120),
					score.AccuracyGTE(0.95),
					score.Or(
						score.KeystrokesGT(userMaxScore.Keystrokes),
						score.And(
							score.KeystrokesEQ(userMaxScore.Keystrokes),
							score.AccuracyGT(userMaxScore.Accuracy),
						),
					),
				),
			).
			Count(ctx)
	default:
		return 0, fmt.Errorf("invalid sort by parameter: %s", sortBy)
	}

	if err != nil {
		return 0, err
	}

	// ランキングを返す
	return rank + 1, nil
}
