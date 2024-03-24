package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func GetScoresRanking(ctx context.Context, sortBy string, start int) ([]*model.ScoreRanking, error) {
	client := ent.FromContext(ctx)

	var rankings []*model.ScoreRanking

	err := client.Score.Query().
		WithUser().
		Order(ent.Desc(sortBy)).
		Limit(50).
		Offset(start-1).
		Select("id, keystrokes, accuracy, created_at").
		Scan(ctx, func(s *ent.Score) {
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

			ranking := &model.ScoreRanking{
				Rank:  start + len(rankings),
				Score: *score,
			}

			rankings = append(rankings, ranking)
		})

	return rankings, err
}

func CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
	client := ent.FromContext(ctx)

	_, err := client.Score.Create().
		SetKeystrokes(keystrokes).
		SetAccuracy(float64(keystrokes)).
		SetUserID(userID).
		Save(ctx)

	return err
}
