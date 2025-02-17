package repository

import (
	"context"

	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
)

// EntTxManager は ent を用いた TxManager の実装
type EntTxManager struct {
	client *ent_generated.Client
}

// NewEntTxManager は EntTxManager を作成する
func NewEntTxManager(client *ent_generated.Client) *EntTxManager {
	return &EntTxManager{client: client}
}

// Execute はトランザクションを開始し、指定された関数を実行する
func (tm *EntTxManager) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.client.Tx(ctx)
	if err != nil {
		return err
	}

	// トランザクションのスコープで関数を実行
	if err := fn(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
