package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
)

type ScoreRepository interface {
	// GetScores はスコアのランキングを取得する
	// ランキングで有効なスコアは、キーストローク数が120以上、正確性が0.95以上
	// ソート順、開始位置、取得件数を指定して、スコアのランキングを取得する
	GetScores(ctx context.Context, sortBy string, start int, limit int) ([]*model.Score, int, error)

	// GetMaxScores はユーザーの最大スコアを取得する
	// ユーザーIDを指定して、最大スコアを取得する
	// 最大スコアが存在しない場合は、nilを返す
	GetMaxScores(ctx context.Context, userID uuid.UUID) (*model.Score, *model.Score, error)

	// CreateScore は新しいスコアを作成する
	// ユーザーID、キーストローク数、正確性、最大スコアかどうかを指定して、スコアを作成する
	// 以前の最大スコアのフラグを更新する
	CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64, isMaxKeystrokes bool, isMaxAccuracy bool) error
}
