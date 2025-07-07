package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type AllergyUsecase struct {
	repo               interfaces.AllergyRepository
	patientAllergyRepo interfaces.PatientsAllergyRepository
}

func NewAllergyUsecase(repo interfaces.AllergyRepository, patientAllergyRepo interfaces.PatientsAllergyRepository) interfaces.AllergyUsecase {
	return &AllergyUsecase{
		repo:               repo,
		patientAllergyRepo: patientAllergyRepo,
	}
}

// TODO:Дописать нормальную связь аллергий с пациентом через третью таблицу

func (u *AllergyUsecase) AddAllergyToPatient(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError) {
	//TODO implement me
	panic("implement me")
	/*
			allergy := entities.PatientsAllergy{
				PatientID:   patientID,
				AllergyID:   allergyID,
				Description: description,
			}

			createdAllergy, err := u.patientAllergyRepo.AddAllergy(&allergy)
			if err != nil {
				//return entities.PatientsAllergy{}, errors.NewDBError("failed to add allergy to patient", err)
			}

			return *createdAllergy, nil
		}

	*/
}
func (u *AllergyUsecase) GetAllergyByPatientID(patientID uint) ([]entities.Allergy, *errors.AppError) {
	allergies, err := u.patientAllergyRepo.GetPatientAllergiesByID(patientID)
	if err != nil {
		//return nil, errors.NewDBError("failed to get patient allergies", err)
	}
	return allergies, nil
}

func (u UseCases) AddAllergyToPatient(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError) {
	//TODO implement me
	panic("implement me")
}

func (u UseCases) GetAllergyByPatientID(patientID uint) ([]entities.Allergy, *errors.AppError) {
	//TODO implement me
	panic("implement me")
}
