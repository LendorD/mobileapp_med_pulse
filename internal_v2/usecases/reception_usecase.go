package usecases

import (
	"errors"

	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type ReceptionUsecase struct {
	repo interfaces.ReceptionRepository
}

func NewReceptionUsecase(repo interfaces.ReceptionRepository) interfaces.ReceptionUsecase {
	return &ReceptionUsecase{repo: repo}
}

func (u *ReceptionUsecase) Create(input models.CreateReceptionRequest) (entities.Reception, *errors.AppError) {
	reception := entities.Reception{
		DoctorID:        input.DoctorID,
		PatientID:       input.PatientID,
		Date:            input.Date,
		Diagnosis:       input.Diagnosis,
		Recommendations: input.Recommendations,
		IsOut:           input.IsOut,
		Status:          entities.StatusScheduled,
		Address:         input.Address,
	}

	createdReception, err := u.repo.Create(&reception)
	if err != nil {
		return entities.Reception{}, errors.NewDBError("failed to create reception", err)
	}

	return *createdReception, nil
}

func (u *ReceptionUsecase) UpdateStatus(id uint, status entities.ReceptionStatus) (entities.Reception, *errors.AppError) {
	reception, err := u.repo.GetByID(id)
	if err != nil {
		return entities.Reception{}, errors.NewDBError("failed to find reception", err)
	}

	reception.Status = status

	updatedReception, err := u.repo.Update(reception)
	if err != nil {
		return entities.Reception{}, errors.NewDBError("failed to update reception status", err)
	}

	return *updatedReception, nil
}

func (u *ReceptionUsecase) GetByID(id uint) (entities.Reception, *errors.AppError) {
	reception, err := u.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Reception{}, errors.NewNotFoundError("reception not found")
		}
		return entities.Reception{}, errors.NewDBError("failed to get reception", err)
	}
	return *reception, nil
}
