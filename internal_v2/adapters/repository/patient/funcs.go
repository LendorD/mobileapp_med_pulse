package patient

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *PatientRepositoryImpl) Create(patient *entities.Patient) error {
	return r.db.Create(patient).Error
}

func (r *PatientRepositoryImpl) Update(patient *entities.Patient) error {
	return r.db.Save(patient).Error
}

func (r *PatientRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Patient{}, id).Error
}

func (r *PatientRepositoryImpl) GetByID(id uint) (*entities.Patient, error) {
	var patient entities.Patient
	if err := r.db.Preload("PersonalInfo").Preload("ContactInfo").First(&patient, id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepositoryImpl) GetAll() ([]entities.Patient, error) {
	var patients []entities.Patient
	if err := r.db.Preload("PersonalInfo").Preload("ContactInfo").Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *PatientRepositoryImpl) GetByFullName(name string) ([]entities.Patient, error) {
	var patients []entities.Patient
	if err := r.db.Where("full_name ILIKE ?", "%"+name+"%").Preload("PersonalInfo").Preload("ContactInfo").Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
