package allergy

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *AllergyRepositoryImpl) CreateAllergy(allergy *entities.Allergy) error {
	return r.db.Create(allergy).Error
}

func (r *AllergyRepositoryImpl) UpdateAllergy(allergy *entities.Allergy) error {
	return r.db.Save(allergy).Error
}

func (r *AllergyRepositoryImpl) DeleteAllergy(id uint) error {
	return r.db.Delete(&entities.Allergy{}, id).Error
}

func (r *AllergyRepositoryImpl) GetAllergyByID(id uint) (*entities.Allergy, error) {
	var a entities.Allergy
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AllergyRepositoryImpl) GetAllergyByName(name string) (*entities.Allergy, error) {
	var a entities.Allergy
	if err := r.db.Where("name = ?", name).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AllergyRepositoryImpl) GetAllAllergy() ([]entities.Allergy, error) {
	var list []entities.Allergy
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
