package receptionSmp

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type ReceptionSmpRepositoryImpl struct {
	db *base.BaseRepository
}

func NewReceptionSmpRepository(db *gorm.DB) interfaces.ReceptionSmpRepository {
	return &ReceptionSmpRepositoryImpl{db: base.NewBaseRepository(db)}
}
