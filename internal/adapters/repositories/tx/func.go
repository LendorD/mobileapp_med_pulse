package tx

import (
	"gorm.io/gorm"
)

func (r *TxRepositoryImpl) BeginTx() (*gorm.DB, error) {
	tx := r.db.Begin()
	return tx, tx.Error
}

func (r *TxRepositoryImpl) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *TxRepositoryImpl) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}
