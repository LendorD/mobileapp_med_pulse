package emergencyReceptionRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type EmergencyReceptionRepositoryImpl struct {
	db *gorm.DB
}

func NewEmergencyReceptionRepository(db *gorm.DB) interfaces.EmergencyReceptionRepository {
	return &EmergencyReceptionRepositoryImpl{db: db}
}
