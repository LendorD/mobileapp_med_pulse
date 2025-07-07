package usecases

import "github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"

type AllergyUsecase struct {
	repo               interfaces.AllergyRepository
	patientAllergyRepo interfaces.PatientAllergyRepository
}

func NewAllergyUsecase(
	repo interfaces.AllergyRepository,
	patientAllergyRepo interfaces.PatientAllergyRepository,
) interfaces.AllergyUsecase {
	return &AllergyUsecase{
		repo:               repo,
		patientAllergyRepo: patientAllergyRepo,
	}
}

func (u *AllergyUsecase) AddToPatient(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError) {
	allergy := entities.PatientsAllergy{
		PatientID:   patientID,
		AllergyID:   allergyID,
		Description: description,
	}

	createdAllergy, err := u.patientAllergyRepo.AddAllergy(&allergy)
	if err != nil {
		return entities.PatientsAllergy{}, errors.NewDBError("failed to add allergy to patient", err)
	}

	return *createdAllergy, nil
}

func (u *AllergyUsecase) GetByPatientID(patientID uint) ([]entities.PatientsAllergy, *errors.AppError) {
	allergies, err := u.patientAllergyRepo.GetPatientAllergies(patientID)
	if err != nil {
		return nil, errors.NewDBError("failed to get patient allergies", err)
	}
	return allergies, nil
}
