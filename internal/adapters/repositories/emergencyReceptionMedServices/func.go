package emergencyReceptionMedServicesRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *EmergencyReceptionMedServicesRepositoryImpl) CreateEmergencyReceptionMedServices(link *entities.EmergencyReceptionMedServices) error {
	return r.db.Create(link).Error
}

func (r *EmergencyReceptionMedServicesRepositoryImpl) DeleteEmergencyReceptionMedServices(id uint) error {
	return r.db.Delete(&entities.EmergencyReceptionMedServices{}, id).Error
}

func (r *EmergencyReceptionMedServicesRepositoryImpl) GetEmergencyReceptionMedServicesByEmergencyReceptionID(erID uint) ([]entities.EmergencyReceptionMedServices, error) {
	var list []entities.EmergencyReceptionMedServices
	err := r.db.Where("emergency_reception_id = ?", erID).Find(&list).Error
	return list, err
}
