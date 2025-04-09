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

// GetScores は指定されたキーストローク数と精度を持つスコアを取得する
func (r *EntScoreRepository) GetScores(ctx context.Context, keystrokes *int, accuracy *float64, sortBy *string) ([]*model.Score, error) {
	query := r.client.Score.Query().WithUser()

	// 条件が指定されている場合のみフィルタを適用
	if keystrokes != nil {
		query = query.Where(score.KeystrokesGTE(*keystrokes))
	}
	if accuracy != nil {
		query = query.Where(score.AccuracyGTE(*accuracy))
	}

	// ソート順が指定されている場合のみ適用
	if sortBy != nil {
		query = query.Order(ent_generated.Desc(*sortBy))
	}

	scores, err := query.All(ctx)
	if err != nil {
		return nil, err
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

	return scoreList, nil
}

// CreateScore は新しいスコアを作成する
func (r *EntScoreRepository) CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
	// スコアを保存
	_, err := r.client.Score.Create().
		SetUserID(userID).
		SetKeystrokes(keystrokes).
		SetAccuracy(accuracy).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
