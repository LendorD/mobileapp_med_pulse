package usecases

import (
	"context"
	"fmt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

type EmkUsecase struct {
	repo       interfaces.EmkRepository
	onecClient interfaces.OneCClient
}

func NewEmkUsecse(repo interfaces.EmkRepository, onecClient interfaces.OneCClient) interfaces.EmkUsecase {
	return &EmkUsecase{
		repo:       repo,
		onecClient: onecClient,
	}
}

// GetAllEmkByPatientID — получает список ЭМК из БД или запрашивает из 1С
func (u *EmkUsecase) GetAllEmkByPatientID(ctx context.Context, patientID string) ([]entities.Emk, error) {
	// 1. Пробуем взять из БД
	emkList, err := u.repo.GetEmkByPatientID(ctx, patientID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("db error: %w", err)
	}
	if len(emkList) > 0 {
		return emkList, nil
	}

	// 2. Если нет в БД — запрашиваем из 1С
	onecEmkList, err := u.onecClient.GetEmkByPatientID(patientID)
	if err != nil {
		return nil, fmt.Errorf("onec error: %w", err)
	}

	// 3. Пробуем сохранить в БД (если упадёт — просто логируем)
	if saveErr := u.repo.SaveEmkList(ctx, patientID, onecEmkList); saveErr != nil {
		fmt.Printf("warn: failed to save EMK for patient %s: %v\n", patientID, saveErr)
	}

	return onecEmkList, nil
}
