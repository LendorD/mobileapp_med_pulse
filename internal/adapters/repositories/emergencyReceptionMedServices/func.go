package emergencyReceptionMedServices

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"gorm.io/gorm"
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

func (r *EmergencyReceptionMedServicesRepositoryImpl) AddService(service *entities.EmergencyReceptionMedServices) (*entities.EmergencyReceptionMedServices, error) {
	// Проверяем, существует ли уже такая связь
	var existingService entities.EmergencyReceptionMedServices
	err := r.db.Where("emergency_reception_id = ? AND med_service_id = ?",
		service.EmergencyReceptionID, service.MedServiceID).
		First(&existingService).Error

	if err == nil {
		// Связь уже существует, возвращаем существующую запись
		return &existingService, nil
	} else if err != gorm.ErrRecordNotFound {
		// Произошла ошибка, отличная от "запись не найдена"
		return nil, err
	}

	// Создаем новую связь
	if err := r.db.Create(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (r *EmergencyReceptionMedServicesRepositoryImpl) GetServicesForEmergency(emergencyID uint) ([]entities.MedService, error) {
	var medServices []entities.MedService

	// Выполняем JOIN через таблицу связи
	err := r.db.
		Model(&entities.EmergencyReceptionMedServices{}).
		Select("med_services.*").
		Joins("JOIN med_services ON med_services.id = emergency_reception_med_services.med_service_id").
		Where("emergency_reception_med_services.emergency_reception_id = ?", emergencyID).
		Find(&medServices).Error

	if err != nil {
		return nil, err
	}

	return medServices, nil
}

// func (r *EmergencyReceptionMedServicesRepositoryImpl) GetServicesForEmergency_(emergencyID uint) ([]entities.MedService, error) {
// 	var emergencyReception entities.EmergencyReception

// 	err := r.db.
// 		Preload("MedServices").
// 		First(&emergencyReception, emergencyID).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return []entities.MedService{}, nil
// 		}
// 		return nil, err
// 	}

// 	return emergencyReception.MedServices, nil
// }
