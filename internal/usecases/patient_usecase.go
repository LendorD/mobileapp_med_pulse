package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type PatientUsecase struct {
	repo interfaces.PatientRepository
}

func NewPatientUsecase(repo interfaces.PatientRepository) interfaces.PatientUsecase {
	return &PatientUsecase{repo: repo}
}

func (u *PatientUsecase) CreatPatient(input *models.CreatePatientRequest) (entities.Patient, *errors.AppError) {
	patient := entities.Patient{
		FullName:  input.FullName,
		BirthDate: input.BirthDate,
		IsMale:    input.IsMale,
	}

	createdPatientId, err := u.repo.CreatePatient(patient)
	if err != nil {
		return entities.Patient{}, errors.NewAppError(errors.InternalServerErrorCode, errors.InternalServerError, err, false)
	}

	createdPatient, err := u.repo.GetPatientByID(createdPatientId)
	if err != nil {
		return entities.Patient{}, errors.NewAppError(errors.InternalServerErrorCode, errors.InternalServerError, err, false)
	}

	return createdPatient, nil
}

func (u *PatientUsecase) GetPatientByID(id uint) (entities.Patient, *errors.AppError) {

	patient, err := u.repo.GetPatientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Patient{}, errors.NewNotFoundError("patient not found")
		}
		return entities.Patient{}, errors.NewAppError(errors.InternalServerErrorCode, errors.InternalServerError, err, false)
	}
	return patient, nil
}

func (u *PatientUsecase) UpdatePatient(input *models.UpdatePatientRequest) (entities.Patient, *errors.AppError) {
	updateMap := map[string]interface{}{
		"id":         input.ID,        // Может быть nil
		"birthdate":  input.BirthDate, // Может быть nil
		"fullname":   input.FullName,  // Может быть nil
		"updated_at": time.Now(),      // Всегда обновляем
	}

	updatedPatientId, err := u.repo.UpdatePatient(input.ID, updateMap)
	if err != nil {
		return entities.Patient{}, errors.NewAppError(errors.InternalServerErrorCode, errors.InternalServerError, err, false)
	}

	updatedPatient, err := u.repo.GetPatientByID(updatedPatientId)
	if err != nil {
		return entities.Patient{}, errors.NewAppError(errors.InternalServerErrorCode, errors.InternalServerError, err, false)
	}

	return updatedPatient, nil

}

func (u *PatientUsecase) DeletePatient(id uint) *errors.AppError {
	if err := u.repo.DeletePatient(id); err.Err != nil {
		return err
	}
	return nil
}
