package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated/predicate"
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
	// 基本のWhere条件
	conditions := []predicate.Score{
		score.KeystrokesGTE(120),
		score.AccuracyGTE(0.95),
	}

	// sortByの値に応じて追加条件を設定
	switch sortBy {
	case "accuracy":
		conditions = append(conditions, score.IsMaxAccuracyEQ(true))
	case "keystrokes":
		conditions = append(conditions, score.IsMaxKeystrokesEQ(true))
	}

	query := r.client.Score.Query().
		WithUser().
		Where(score.And(conditions...)).
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
			User: model.User{
				ID:            s.Edges.User.ID.String(),
				StudentNumber: s.Edges.User.StudentNumber,
				HandleName:    s.Edges.User.HandleName,
				CreatedAt:     s.Edges.User.CreatedAt,
				UpdatedAt:     s.Edges.User.UpdatedAt,
			},
		}
	}

	return scoreList, totalCount, nil
}

// GetMaxScores はユーザーの現在の最大スコアを取得する
func (r *EntScoreRepository) GetMaxScores(ctx context.Context, userID uuid.UUID) (*model.Score, *model.Score, error) {
	maxKeystrokeScore, err := r.client.Score.Query().
		WithUser().
		Where(score.UserID(userID), score.IsMaxKeystrokes(true)).
		// TODO: isMaxKeyStorokesとかがUniqueじゃないのが悪さしている https://github.com/su-its/typing/issues/203
		First(ctx) // NOTE: バグでUniqueで検索するとエラーになるため、Firstで検索する
	if err != nil && !ent_generated.IsNotFound(err) {
		return nil, nil, err
	}

	maxAccuracyScore, err := r.client.Score.Query().
		WithUser().
		Where(score.UserID(userID), score.IsMaxAccuracy(true)).
		// TODO: isMaxAccuracyとかがUniqueじゃないのが悪さしている https://github.com/su-its/typing/issues/203
		First(ctx) // NOTE: バグでUniqueで検索するとエラーになるため、Firstで検索する
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
			User: model.User{
				ID:            maxKeystrokeScore.Edges.User.ID.String(),
				StudentNumber: maxKeystrokeScore.Edges.User.StudentNumber,
				HandleName:    maxKeystrokeScore.Edges.User.HandleName,
				CreatedAt:     maxKeystrokeScore.Edges.User.CreatedAt,
				UpdatedAt:     maxKeystrokeScore.Edges.User.UpdatedAt,
			},
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
			User: model.User{
				ID:            maxAccuracyScore.Edges.User.ID.String(),
				StudentNumber: maxAccuracyScore.Edges.User.StudentNumber,
				HandleName:    maxAccuracyScore.Edges.User.HandleName,
				CreatedAt:     maxAccuracyScore.Edges.User.CreatedAt,
				UpdatedAt:     maxAccuracyScore.Edges.User.UpdatedAt,
			},
		}
	}

	return maxKeystrokeScoreModel, maxAccuracyScoreModel, nil
}

// CreateScore は新しいスコアを作成する
func (r *EntScoreRepository) CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64, isMaxKeystrokes bool, isMaxAccuracy bool) error {
	// 以前の最大スコアのフラグを更新
	if isMaxKeystrokes {
		err := r.client.Score.Update().
			Where(score.And(score.UserID(userID), score.IsMaxKeystrokes(true))).
			SetIsMaxKeystrokes(false).
			Exec(ctx)
		if err != nil && !ent_generated.IsNotFound(err) {
			return err
		}
	}

	if isMaxAccuracy {
		err := r.client.Score.Update().
			Where(score.And(score.UserID(userID), score.IsMaxAccuracy(true))).
			SetIsMaxAccuracy(false).
			Exec(ctx)
		if err != nil && !ent_generated.IsNotFound(err) {
			return err
		}
	}

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

	return nil
}
