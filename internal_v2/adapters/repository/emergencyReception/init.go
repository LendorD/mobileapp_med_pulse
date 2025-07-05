package emergencyReception

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type EmergencyReceptionRepositoryImpl struct {
	db *gorm.DB
}

func NewEmergencyReceptionRepository(db *gorm.DB) interfaces.EmergencyReceptionRepository {
	return &EmergencyReceptionRepositoryImpl{db: db}
}
