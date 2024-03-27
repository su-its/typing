package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/score"
)

func GetScoresRanking(ctx context.Context, client *ent.Client, sortBy string, start, limit int) ([]*model.ScoreRanking, error) {
	// ! このクエリは要チェック
	scores, err := client.Score.Query().
		WithUser().
		Where(
			score.And(
				score.KeystrokesGTE(120),
				score.AccuracyGTE(95.0),
			),
		).
		Order(ent.Desc(sortBy)).
		Limit(limit).
		Offset(start-1).
		Select("user_id").
		Unique(true).
		Select("id", "keystrokes", "accuracy", "created_at", "user_id").
		All(ctx)
	if err != nil {
		return nil, err
	}

	var rankings []*model.ScoreRanking
	var prevScore float64
	var rank int

	for i, s := range scores {
		user := &model.User{
			ID:            s.Edges.User.ID.String(),
			StudentNumber: s.Edges.User.StudentNumber,
			HandleName:    s.Edges.User.HandleName,
			CreatedAt:     s.Edges.User.CreatedAt,
			UpdatedAt:     s.Edges.User.UpdatedAt,
		}

		score := &model.Score{
			ID:         s.ID.String(),
			Keystrokes: s.Keystrokes,
			Accuracy:   s.Accuracy,
			CreatedAt:  s.CreatedAt,
			User:       *user,
		}

		var currentScore float64
		switch sortBy {
		case "accuracy":
			currentScore = s.Accuracy
		case "keystrokes":
			currentScore = float64(s.Keystrokes)
		default:
			return nil, fmt.Errorf("invalid sort by parameter: %s", sortBy)
		}

		if i == 0 || currentScore != prevScore {
			rank = start + i
		}
		prevScore = currentScore

		ranking := &model.ScoreRanking{
			Rank:  rank,
			Score: *score,
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

func CreateScore(ctx context.Context, client *ent.Client, userID uuid.UUID, keystrokes int, accuracy float64) error {
	_, err := client.Score.Create().
		SetKeystrokes(keystrokes).
		SetAccuracy(float64(keystrokes)).
		SetUserID(userID).
		Save(ctx)

	return err
}
