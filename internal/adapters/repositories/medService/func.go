package medServiceRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *MedServiceRepositoryImpl) CreateMedService(service *entities.MedService) error {
	return r.db.Create(service).Error
}

func (r *MedServiceRepositoryImpl) UpdateMedService(service *entities.MedService) error {
	return r.db.Save(service).Error
}

func (r *MedServiceRepositoryImpl) DeleteMedService(id uint) error {
	return r.db.Delete(&entities.MedService{}, id).Error
}

func (r *MedServiceRepositoryImpl) GetMedServiceByID(id uint) (*entities.MedService, error) {
	var s entities.MedService
	if err := r.db.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *MedServiceRepositoryImpl) GetMedServiceByName(name string) (*entities.MedService, error) {
	var s entities.MedService
	if err := r.db.Where("name = ?", name).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *MedServiceRepositoryImpl) GetAllMedService() ([]entities.MedService, error) {
	var list []entities.MedService
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
