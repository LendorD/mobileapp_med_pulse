package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	_ "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	_ "github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type UseCases struct {
	interfaces.AllergyUsecase
	interfaces.ContactInfoUsecase
	interfaces.DoctorUsecase
	interfaces.EmergencyReceptionUsecase
	interfaces.EmergencyReceptionMedServicesUsecase
	interfaces.MedServiceUsecase
	interfaces.PatientUsecase
	interfaces.PersonalInfoUsecase
	interfaces.ReceptionUsecase
}

func NewUsecases(r interfaces.Repository, s interfaces.Service, conf *config.Config) interfaces.Usecases {

	// Создание структуры UseCases
	return &UseCases{}

}
