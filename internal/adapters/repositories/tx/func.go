package tx

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"gorm.io/gorm"
)

// Begin начинает транзакцию и кладёт её в контекст
func (tm *TxManager) Begin(ctx context.Context) (context.Context, error) {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return context.WithValue(ctx, base.TxContextKey, tx), nil
}

func (tm *TxManager) Commit(ctx context.Context) error {
	tx := tm.GetTransaction(ctx)
	if tx == nil {
		return nil // нет транзакции — ничего не делаем
	}
	return tx.Commit().Error
}

// Rollback откатывает транзакцию из контекста
func (tm *TxManager) Rollback(ctx context.Context) error {
	tx := tm.GetTransaction(ctx)
	if tx == nil {
		return nil
	}
	return tx.Rollback().Error
}

// getTransaction извлекает транзакцию из контекста
func (tm *TxManager) GetTransaction(ctx context.Context) *gorm.DB {
	tx, _ := ctx.Value(base.TxContextKey).(*gorm.DB)
	return tx
}
