package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type MedServiceUsecase struct {
	repo interfaces.MedServiceRepository
}

func NewMedServiceUsecase(repo interfaces.MedServiceRepository) interfaces.MedServiceUsecase {
	return &MedServiceUsecase{
		repo: repo,
	}
}

func (u *MedServiceUsecase) GetAllMedServices() (models.MedServicesListResponse, *errors.AppError) {
	empty := models.MedServicesListResponse{}

	services, totalRows, err := u.repo.GetAllMedServices()
	if err != nil {
		return empty, errors.NewAppError(errors.InternalServerErrorCode, "failed to get medical services", err, true)
	}

	var result []models.MedServicesResponse
	for _, s := range services {
		result = append(result, models.MedServicesResponse{
			ID:    s.ID,
			Name:  s.Name,
			Price: s.Price,
		})
	}

	return models.MedServicesListResponse{
		Hits:      result,
		TotalHits: int(totalRows),
	}, nil
}
