package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
)

type ScoreRepository interface {
	// GetScores は指定されたキーストローク数と精度を持つスコアを取得する
	// sortByには"accuracy"か"keystrokes"を指定する
	// TODO: optinal
	GetScores(ctx context.Context, keystrokes int, accuracy float64, sortBy string) ([]*model.Score, error)

	// GetMaxScores はユーザーの最大スコアを取得する
	// ユーザーIDを指定して、最大スコアを取得する
	// 最大スコアが存在しない場合は、nilを返す
	GetMaxScores(ctx context.Context, userID uuid.UUID) (*model.Score, *model.Score, error)

	// CreateScore は新しいスコアを作成する
	// ユーザーID、キーストローク数、正確性、最大スコアかどうかを指定して、スコアを作成する
	// 以前の最大スコアのフラグを更新する
	CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64, isMaxKeystrokes bool, isMaxAccuracy bool) error
}
