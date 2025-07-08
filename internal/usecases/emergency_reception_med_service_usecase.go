package usecases

import "github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"

type EmergencyReceptionMedServicesUsecase struct {
	repo interfaces.EmergencyReceptionMedServicesRepository
}

func NewEmergencyReceptionMedServicesUsecase(repo interfaces.EmergencyReceptionMedServicesRepository) interfaces.EmergencyReceptionMedServicesUsecase {
	return &EmergencyReceptionMedServicesUsecase{repo: repo}
}
