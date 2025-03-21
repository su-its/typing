package repository

import "context"

// TxManager はトランザクションを管理するインターフェース
type TxManager interface {
	// Execute はトランザクションを開始し、指定した処理を実行する
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}
