package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type UseCases struct {
	interfaces.SmpService
	interfaces.ReceptionService
	interfaces.PatientService
}

func NewUsecases(r interfaces.Repository, s interfaces.Service, conf *config.Config) interfaces.Usecases {

	// Создание структуры UseCases
	return &UseCases{}

}
