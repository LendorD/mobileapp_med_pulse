package emergencyReceptionMedServicesRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *EmergencyReceptionMedServicesRepositoryImpl) Create(link *entities.EmergencyReceptionMedServices) error {
	return r.db.Create(link).Error
}

func (r *EmergencyReceptionMedServicesRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.EmergencyReceptionMedServices{}, id).Error
}

func (r *EmergencyReceptionMedServicesRepositoryImpl) GetByEmergencyReceptionID(erID uint) ([]entities.EmergencyReceptionMedServices, error) {
	var list []entities.EmergencyReceptionMedServices
	err := r.db.Where("emergency_reception_id = ?", erID).Find(&list).Error
	return list, err
}
