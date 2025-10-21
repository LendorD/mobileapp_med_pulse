package medcard

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type MedicalCardRepository struct {
	db *base.BaseRepository
}

func NewMedicalCardRepository(db *gorm.DB) interfaces.MedicalCardRepository {
	return &MedicalCardRepository{db: base.NewBaseRepository(db)}
}
