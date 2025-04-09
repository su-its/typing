package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
)

type ScoreRepository interface {
	// GetScores は条件に合うスコアを取得する
	// sortByには"accuracy"か"keystrokes"を指定する
	// keystrokes, accuracy, sortByはすべてオプショナル
	GetScores(ctx context.Context, keystrokes *int, accuracy *float64, sortBy *string) ([]*model.Score, error)

	// CreateScore は新しいスコアを作成する
	CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error
}
