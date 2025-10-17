package medcard

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type MedicalCardRepository struct {
	db *gorm.DB
}

func NewMedicalCardRepository(db *gorm.DB) interfaces.MedicalCardRepository {
	return &MedicalCardRepository{db: db}
}
