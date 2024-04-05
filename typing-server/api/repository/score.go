package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/model"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
	"github.com/su-its/typing/typing-server/domain/repository/ent/score"
	"github.com/su-its/typing/typing-server/domain/repository/ent/user"
)

func GetScoresRanking(ctx context.Context, client *ent.Client, sortBy string, start, limit int) ([]*model.ScoreRanking, int, error) {
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
		return nil, 0, fmt.Errorf("invalid sort by parameter: %s", sortBy)
	}

	//全件数の取得
	count := query.CountX(ctx)

	// フラグでフィルタリングされたスコアを取得
	scores, err := query.
		Limit(limit).
		Offset(start - 1).
		All(ctx)

	if err != nil {
		return nil, 0, err
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

	return rankings, count, nil
}
func CreateScore(ctx context.Context, client *ent.Client, userID uuid.UUID, keystrokes int, accuracy float64) error {
	// トランザクションを開始
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		// txがコミット済みの場合はロールバックしない
		if tx == nil {
			return
		}
		if err := tx.Rollback(); err != nil {
			if err != sql.ErrTxDone {
				return
			}
			return
		}
	}()

	// 新しいスコアを作成
	createdScore, err := tx.Score.Create().
		SetKeystrokes(keystrokes).
		SetAccuracy(accuracy).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return err
	}

	// キーストロークと精度のチェック
	if keystrokes < 120 || accuracy < 0.95 {
		return tx.Commit()
	}

	// ユーザーを取得
	user, err := tx.User.Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		return err
	}

	// 最大キーストロークと最大精度のスコアを取得
	maxKeystrokeScore, maxAccuracyScore, err := getMaxScores(ctx, user)
	if err != nil {
		return err
	}

	// 以前の最大スコアのフラグを更新
	err = updateMaxScoreFlags(ctx, maxKeystrokeScore, maxAccuracyScore, createdScore)
	if err != nil {
		return err
	}

	// トランザクションをコミット
	return tx.Commit()
}

func getMaxScores(ctx context.Context, user *ent.User) (*ent.Score, *ent.Score, error) {
	// 最大キーストロークのスコアを取得
	maxKeystrokeScore, err := user.QueryScores().Where(score.IsMaxKeystrokes(true)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, nil, err
	}

	// 最大精度のスコアを取得
	maxAccuracyScore, err := user.QueryScores().Where(score.IsMaxAccuracy(true)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, nil, err
	}

	return maxKeystrokeScore, maxAccuracyScore, nil
}

func updateMaxScoreFlags(ctx context.Context, maxKeystrokeScore, maxAccuracyScore, createdScore *ent.Score) error {
	isMaxKeystrokes := maxKeystrokeScore == nil || createdScore.Keystrokes >= maxKeystrokeScore.Keystrokes
	isMaxAccuracy := maxAccuracyScore == nil || createdScore.Accuracy >= maxAccuracyScore.Accuracy

	// 以前の最大スコアのフラグを更新
	if maxKeystrokeScore != nil && isMaxKeystrokes {
		err := maxKeystrokeScore.Update().SetIsMaxKeystrokes(false).Exec(ctx)
		if err != nil {
			return err
		}
	}

	if maxAccuracyScore != nil && isMaxAccuracy {
		err := maxAccuracyScore.Update().SetIsMaxAccuracy(false).Exec(ctx)
		if err != nil {
			return err
		}
	}

	// 作成したスコアの最大フラグを更新
	err := createdScore.Update().
		SetIsMaxKeystrokes(isMaxKeystrokes).
		SetIsMaxAccuracy(isMaxAccuracy).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
