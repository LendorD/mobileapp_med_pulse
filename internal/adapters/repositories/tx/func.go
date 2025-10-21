package tx

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const txContextKey contextKey = "db_transaction"

func (tm *TxManager) Commit(ctx context.Context) error {
	tx := tm.GetTransaction(ctx)
	if tx == nil {
		return nil
	}
	return tx.Commit().Error
}

func (tm *TxManager) Rollback(ctx context.Context) error {
	tx := tm.GetTransaction(ctx)
	if tx == nil {
		return nil
	}
	return tx.Rollback().Error
}

func (tm *TxManager) GetTransaction(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txContextKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}
