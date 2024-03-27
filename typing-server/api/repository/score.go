package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/score"
	"github.com/su-its/typing/typing-server/domain/repository/ent/user"
)

func GetScoresRanking(ctx context.Context, client *ent.Client, sortBy string, start, limit int) ([]*model.ScoreRanking, error) {
	var scores []*ent.Score

	// entのクエリを使用してスコアを取得
	query := client.Score.Query().
		WithUser().
		Where(
			score.And(
				score.KeystrokesGTE(120),
				score.AccuracyGTE(0.95),
			),
		).
		Order(ent.Desc(sortBy))

	switch sortBy {
	case "accuracy":
		query = query.Where(score.IsMaxAccuracy(true))
	case "keystrokes":
		query = query.Where(score.IsMaxKeystrokes(true))
	default:
		return nil, fmt.Errorf("invalid sort by parameter: %s", sortBy)
	}

	// フラグでフィルタリングされたスコアを取得
	scores, err := query.
		Limit(limit).
		Offset(start - 1).
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
			UserID:     s.UserID.String(),
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
	// Create a new score
	createdScore, err := client.Score.Create().
		SetKeystrokes(keystrokes).
		SetAccuracy(accuracy).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return err
	}

	if keystrokes < 120 || accuracy < 0.95 {
		return nil
	}

	// Get the user
	user, err := client.User.Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		return err
	}

	// Check if the new score has the maximum keystrokes
	maxKeystrokeScore, err := user.QueryScores().Where(score.IsMaxKeystrokes(true)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	isMaxKeystrokes := maxKeystrokeScore == nil || createdScore.Keystrokes >= maxKeystrokeScore.Keystrokes

	// Check if the new score has the maximum accuracy
	maxAccuracyScore, err := user.QueryScores().Where(score.IsMaxAccuracy(true)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	isMaxAccuracy := maxAccuracyScore == nil || createdScore.Accuracy >= maxAccuracyScore.Accuracy

	// Update the flags of the previous maximum scores
	if maxKeystrokeScore != nil && !isMaxKeystrokes {
		err = maxKeystrokeScore.Update().SetIsMaxKeystrokes(false).Exec(ctx)
		if err != nil {
			return err
		}
	}
	if maxAccuracyScore != nil && !isMaxAccuracy {
		err = maxAccuracyScore.Update().SetIsMaxAccuracy(false).Exec(ctx)
		if err != nil {
			return err
		}
	}

	// Update the score with the maximum flags
	err = createdScore.Update().
		SetIsMaxKeystrokes(isMaxKeystrokes).
		SetIsMaxAccuracy(isMaxAccuracy).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
