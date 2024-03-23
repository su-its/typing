package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func GetScoresRanking(ctx context.Context, sortBy string, start, limit int) ([]*model.ScoreRanking, error) {
	client := ent.FromContext(ctx)

	var rankings []*model.ScoreRanking

	err := client.Score.Query().
		WithUser().
		Order(ent.Desc(sortBy)).
		Limit(limit).
		Offset(start-1).
		Select("id, keystrokes, accuracy, created_at").
		Scan(ctx, func(s *ent.Score) {
			rankings = append(rankings, &model.ScoreRanking{
				Rank:      start + len(rankings),
				UserID:    s.ID.String(),
				Username:  s.Edges.User.HandleName,
				Score:     s.Keystrokes,
				Accuracy:  s.Accuracy,
				CreatedAt: s.CreatedAt,
			})
		})
	return rankings, err
}

func CreateScore(ctx context.Context, score *ent.Score) error {
	client := ent.FromContext(ctx)

	_, err := client.Score.Create().
		SetKeystrokes(score.Keystrokes).
		SetAccuracy(score.Accuracy).
		SetUser(score.Edges.User).
		Save(ctx)

	return err
}
