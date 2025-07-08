package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type ReceptionUsecase struct {
	repo interfaces.ReceptionRepository
}

func NewReceptionUsecase(repo interfaces.ReceptionRepository) interfaces.ReceptionUsecase {
	return &ReceptionUsecase{repo: repo}
}

// func (u *ReceptionUsecase) Create(input models.CreateReceptionRequest) (entities.Reception, *errors.AppError) {
// 	reception := entities.Reception{
// 		DoctorID:        input.DoctorID,
// 		PatientID:       input.PatientID,
// 		Date:            input.Date,
// 		Diagnosis:       input.Diagnosis,
// 		Recommendations: input.Recommendations,
// 		IsOut:           input.IsOut,
// 		Status:          entities.StatusScheduled,
// 		Address:         input.Address,
// 	}

// 	createdReception, err := u.repo.Create(&reception)
// 	if err != nil {
// 		return entities.Reception{}, errors.NewDBError("failed to create reception", err)
// 	}

// 	return *createdReception, nil
// }

// func (u *ReceptionUsecase) UpdateStatus(id uint, status entities.ReceptionStatus) (entities.Reception, *errors.AppError) {
// 	reception, err := u.repo.GetByID(id)
// 	if err != nil {
// 		return entities.Reception{}, errors.NewDBError("failed to find reception", err)
// 	}

// 	reception.Status = status

// 	updatedReception, err := u.repo.Update(reception)
// 	if err != nil {
// 		return entities.Reception{}, errors.NewDBError("failed to update reception status", err)
// 	}

// 	return *updatedReception, nil
// }

// func (u *ReceptionUsecase) GetByID(id uint) (entities.Reception, *errors.AppError) {
// 	reception, err := u.repo.GetByID(id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return entities.Reception{}, errors.NewNotFoundError("reception not found")
// 		}
// 		return entities.Reception{}, errors.NewDBError("failed to get reception", err)
// 	}
// 	return *reception, nil
// }

// isTimeOverlap проверяет пересечение временных интервалов
func isTimeOverlap(t1, t2 time.Time, duration time.Duration) bool {
	end1 := t1.Add(duration)
	end2 := t2.Add(duration)
	return t1.Before(end2) && end1.After(t2)
}

// Обновленный метод сервиса
func (s *ReceptionUsecase) GetReceptionsByDoctorAndDate(doctorID uint, date time.Time, page int) ([]models.ReceptionShortResponse, error) {

	// Валидация номера страницы
	if page < 1 {
		return nil, errors.New("page must be greater than 0")
	}

	// Количество записей на странице
	const perPage = 3
	// Получаем данные из репозитория
	receptions, err := s.repo.GetReceptionsByDoctorAndDate(doctorID, date, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to get receptions: %w", err)
	}

	return receptions, nil
}
