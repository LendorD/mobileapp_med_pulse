package base

import (
	"context"
	"log"

	"gorm.io/gorm"
)

// GetDB возвращает транзакцию из контекста или основное подключение
func (br *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TxContextKey).(*gorm.DB); ok && tx != nil {
		log.Println("✅ Using transaction DB")
		return tx
	}
	log.Println("⚠️ Using main DB (NO TRANSACTION!)")
	return br.db
}
