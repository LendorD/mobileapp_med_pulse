package usecases

import "github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"

type EmergencyReceptionMedServicesUsecase struct {
	repo interfaces.EmergencyReceptionMedServicesUsecase
}

func NewEmergencyReceptionMedServicesUsecase(repo interfaces.EmergencyReceptionMedServicesRepository) interfaces.EmergencyReceptionMedServicesUsecase {
	return &EmergencyReceptionMedServicesUsecase{repo: repo}
}
