package emergencyReceptionMedServicesRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type EmergencyReceptionMedServicesRepositoryImpl struct {
	db *gorm.DB
}

func NewEmergencyReceptionMedServicesRepository(db *gorm.DB) interfaces.EmergencyReceptionMedServicesRepository {
	return &EmergencyReceptionMedServicesRepositoryImpl{db: db}
}
