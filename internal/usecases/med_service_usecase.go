package usecases

import "github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"

type MedServiceUsecase struct {
	repo                    interfaces.MedServiceRepository
	emergencyMedServiceRepo interfaces.EmergencyMedServiceRepository
}

func NewMedServiceUsecase(
	repo interfaces.MedServiceRepository,
	emergencyMedServiceRepo interfaces.EmergencyMedServiceRepository,
) interfaces.MedServiceUsecase {
	return &MedServiceUsecase{
		repo:                    repo,
		emergencyMedServiceRepo: emergencyMedServiceRepo,
	}
}

func (u *MedServiceUsecase) AddToEmergency(emergencyID, serviceID uint) (entities.EmergencyReceptionMedServices, *errors.AppError) {
	service := entities.EmergencyReceptionMedServices{
		EmergencyReceptionID: emergencyID,
		MedServiceID:         serviceID,
	}

	createdService, err := u.emergencyMedServiceRepo.AddService(&service)
	if err != nil {
		return entities.EmergencyReceptionMedServices{}, errors.NewDBError("failed to add service to emergency", err)
	}

	return *createdService, nil
}

func (u *MedServiceUsecase) GetByEmergencyID(emergencyID uint) ([]entities.MedService, *errors.AppError) {
	services, err := u.emergencyMedServiceRepo.GetServicesForEmergency(emergencyID)
	if err != nil {
		return nil, errors.NewDBError("failed to get emergency services", err)
	}
	return services, nil
}
