package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated/score"
)

// EntScoreRepository は ent を使用して ScoreRepository を実装する。
type EntScoreRepository struct {
	client *ent_generated.Client
}

// コンパイル時チェック: EntScoreRepository が ScoreRepository インターフェースを実装していることを保証する。
var _ repository.ScoreRepository = (*EntScoreRepository)(nil)

// NewEntScoreRepository は EntScoreRepository のコンストラクタ。
func NewEntScoreRepository(client *ent_generated.Client) *EntScoreRepository {
	return &EntScoreRepository{client: client}
}

func (r *EntScoreRepository) GetScores(ctx context.Context, sortBy string, start int, limit int) ([]*model.Score, int, error) {
	query := r.client.Score.Query().
		WithUser().
		Where(
			score.And(
				score.KeystrokesGTE(120),
				score.AccuracyGTE(0.95),
			),
		).
		Order(ent_generated.Desc(sortBy))

	totalCount := query.CountX(ctx)

	scores, err := query.Limit(limit).Offset(start - 1).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// ドメインモデルに変換
	scoreList := make([]*model.Score, len(scores))
	for i, s := range scores {
		scoreList[i] = &model.Score{
			ID:         s.ID.String(),
			UserID:     s.UserID.String(),
			Keystrokes: s.Keystrokes,
			Accuracy:   s.Accuracy,
			CreatedAt:  s.CreatedAt,
		}
	}

	return scoreList, totalCount, nil
}

// GetMaxScores はユーザーの現在の最大スコアを取得する
func (r *EntScoreRepository) GetMaxScores(ctx context.Context, userID uuid.UUID) (*model.Score, *model.Score, error) {
	maxKeystrokeScore, err := r.client.Score.Query().
		Where(score.UserID(userID), score.IsMaxKeystrokes(true)).
		Only(ctx)
	if err != nil && !ent_generated.IsNotFound(err) {
		return nil, nil, err
	}

	maxAccuracyScore, err := r.client.Score.Query().
		Where(score.UserID(userID), score.IsMaxAccuracy(true)).
		Only(ctx)
	if err != nil && !ent_generated.IsNotFound(err) {
		return nil, nil, err
	}

	// ドメインモデルに変換
	var maxKeystrokeScoreModel *model.Score
	if maxKeystrokeScore != nil {
		maxKeystrokeScoreModel = &model.Score{
			ID:         maxKeystrokeScore.ID.String(),
			UserID:     maxKeystrokeScore.UserID.String(),
			Keystrokes: maxKeystrokeScore.Keystrokes,
			Accuracy:   maxKeystrokeScore.Accuracy,
			CreatedAt:  maxKeystrokeScore.CreatedAt,
		}
	}

	var maxAccuracyScoreModel *model.Score
	if maxAccuracyScore != nil {
		maxAccuracyScoreModel = &model.Score{
			ID:         maxAccuracyScore.ID.String(),
			UserID:     maxAccuracyScore.UserID.String(),
			Keystrokes: maxAccuracyScore.Keystrokes,
			Accuracy:   maxAccuracyScore.Accuracy,
			CreatedAt:  maxAccuracyScore.CreatedAt,
		}
	}

	return maxKeystrokeScoreModel, maxAccuracyScoreModel, nil
}

// CreateScore は新しいスコアを作成する
func (r *EntScoreRepository) CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64, isMaxKeystrokes bool, isMaxAccuracy bool) error {
	// スコアを保存
	_, err := r.client.Score.Create().
		SetUserID(userID).
		SetKeystrokes(keystrokes).
		SetAccuracy(accuracy).
		SetIsMaxKeystrokes(isMaxKeystrokes).
		SetIsMaxAccuracy(isMaxAccuracy).
		Save(ctx)
	if err != nil {
		return err
	}

	// 以前の最大スコアのフラグを更新
	if isMaxKeystrokes {
		err := r.client.Score.Update().
			Where(score.UserID(userID), score.IsMaxKeystrokes(true)).
			SetIsMaxKeystrokes(false).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	if isMaxAccuracy {
		err := r.client.Score.Update().
			Where(score.UserID(userID), score.IsMaxAccuracy(true)).
			SetIsMaxAccuracy(false).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
