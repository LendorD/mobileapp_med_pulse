package tx

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type TxRepositoryImpl struct {
	db *gorm.DB
}

func NewTxRepository(db *gorm.DB) interfaces.TxRepository {
	return &TxRepositoryImpl{db: db}
}
