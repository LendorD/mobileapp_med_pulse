package patient

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *PatientRepositoryImpl) CreatePatient(patient *entities.Patient) error {
	return r.db.Create(patient).Error
}

func (r *PatientRepositoryImpl) UpdatePatient(patient *entities.Patient) error {
	return r.db.Save(patient).Error
}

func (r *PatientRepositoryImpl) DeletePatient(id uint) error {
	return r.db.Delete(&entities.Patient{}, id).Error
}

func (r *PatientRepositoryImpl) GetPatientByID(id uint) (*entities.Patient, error) {
	var patient entities.Patient
	if err := r.db.Preload("PersonalInfo").Preload("ContactInfo").First(&patient, id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepositoryImpl) GetAllPatient() ([]entities.Patient, error) {
	var patients []entities.Patient
	if err := r.db.Preload("PersonalInfo").Preload("ContactInfo").Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *PatientRepositoryImpl) GetPatientByFullName(name string) ([]entities.Patient, error) {
	var patients []entities.Patient
	if err := r.db.Where("full_name ILIKE ?", "%"+name+"%").Preload("PersonalInfo").Preload("ContactInfo").Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
