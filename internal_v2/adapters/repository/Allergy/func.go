package Allergy

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *AllergyRepositoryImpl) Create(allergy *entities.Allergy) error {
	return r.db.Create(allergy).Error
}

func (r *AllergyRepositoryImpl) Update(allergy *entities.Allergy) error {
	return r.db.Save(allergy).Error
}

func (r *AllergyRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Allergy{}, id).Error
}

func (r *AllergyRepositoryImpl) GetByID(id uint) (*entities.Allergy, error) {
	var a entities.Allergy
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AllergyRepositoryImpl) GetByName(name string) (*entities.Allergy, error) {
	var a entities.Allergy
	if err := r.db.Where("name = ?", name).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AllergyRepositoryImpl) GetAll() ([]entities.Allergy, error) {
	var list []entities.Allergy
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
