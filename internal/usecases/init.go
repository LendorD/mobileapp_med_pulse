package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	_ "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	_ "github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type UseCases struct {
}

func (u UseCases) AddAllergyToPatient(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError) {
	//TODO implement me
	panic("implement me")
}

func (u UseCases) GetAllergyByPatientID(patientID uint) ([]entities.Allergy, *errors.AppError) {
	//TODO implement me
	panic("implement me")
}

func (u UseCases) RemoveAllergyFromPatient(patientID, allergyID uint) *errors.AppError {
	//TODO implement me
	panic("implement me")
}

func (u UseCases) UpdateAllergyDescription(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError) {
	//TODO implement me
	panic("implement me")
}

func NewUsecases(r interfaces.Repository, s interfaces.Service, conf *config.Config) interfaces.Usecases {

	// Создание структуры UseCases
	return &UseCases{}

}
