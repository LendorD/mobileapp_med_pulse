package patient

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm/clause"
)

func (r *PatientRepositoryImpl) CreatePatient(patient entities.Patient) (uint, *errors.AppError) {
	op := "repo.Patient.CreatePatient"

	err := r.db.Clauses(clause.Returning{}).Create(patient).Error
	if err != nil {
		return 0, errors.NewDBError(op, err)
	}
	return patient.ID, nil
}

func (r *PatientRepositoryImpl) UpdatePatient(id uint, updateMap map[string]interface{}) (uint, *errors.AppError) {
	op := "repo.Patient.UpdatePatient"

	var updatedPatient entities.Patient
	result := r.db.
		Clauses(clause.Returning{}).
		Model(&updatedPatient).
		Where("id = ?", id).
		Updates(updateMap)

	if result.Error != nil {
		return 0, errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, errors.NewDBError(op, result.Error)
	}

	return updatedPatient.ID, nil
}

func (r *PatientRepositoryImpl) DeletePatient(id uint) *errors.AppError {
	op := "repo.Patient.DeletePatient"

	result := r.db.Delete(&entities.Patient{}, id)
	if result.Error != nil {
		return errors.NewDBError(op, result.Error)
	}

	return nil
}

func (r *PatientRepositoryImpl) GetPatientByID(id uint) (entities.Patient, *errors.AppError) {
	op := "repo.Patient.GetPatientByID"

	var patient entities.Patient
	err := r.db.Preload("Status").First(&patient, id).Error
	if err != nil {
		return entities.Patient{}, errors.NewDBError(op, err)
	}

	return patient, nil
}

func (r *PatientRepositoryImpl) GetAllPatients() ([]entities.Patient, *errors.AppError) {
	op := "repo.Patient.GetAllPatients"

	var patients []entities.Patient
	if err := r.db.
		Preload("PersonalInfo").
		Preload("ContactInfo").
		Find(&patients).Error; err != nil {

		return nil, errors.NewDBError(op, err)
	}

	if len(patients) == 0 {
		//return nil,  errors.NewDBError(op, error(Nor))
	}

	return patients, nil
}

func (r *PatientRepositoryImpl) GetPatientsByFullName(name string) ([]entities.Patient, *errors.AppError) {
	op := "repo.Patient.GetPatientsByFullName"

	var patients []entities.Patient
	if err := r.db.
		Where("full_name ILIKE ?", "%"+name+"%").
		Preload("PersonalInfo").
		Preload("ContactInfo").
		Find(&patients).Error; err != nil {

		return nil, errors.NewDBError(op, err)
	}

	if len(patients) == 0 {
		//return nil, errors.NewDBError(op, )
	}

	return patients, nil
}
