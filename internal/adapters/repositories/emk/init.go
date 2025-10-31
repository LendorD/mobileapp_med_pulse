package emk

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type EmkRepository struct {
	db *base.BaseRepository
}

func NewEmkRepository(db *gorm.DB) interfaces.EmkRepository {
	return &EmkRepository{
		db: base.NewBaseRepository(db),
	}
}
