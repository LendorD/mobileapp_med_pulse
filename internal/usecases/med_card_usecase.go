package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type MedCardUsecase struct {
	cacheRepo  interfaces.OneCCacheRepository
	onecClient interfaces.OneCClient
}

func NewMedCardUsecase(
	cacheRepo interfaces.OneCCacheRepository,
	onecClient interfaces.OneCClient,
) interfaces.MedCardUsecase {
	return &MedCardUsecase{
		cacheRepo:  cacheRepo,
		onecClient: onecClient,
	}
}

func (u *MedCardUsecase) GetMedCardByPatientID(ctx context.Context, patientID string) (*models.PatientCard, error) {
	// 1. Проверяем кеш
	card, err := u.cacheRepo.GetMedicalCard(ctx, patientID)
	if err != nil {
		return nil, err
	}
	if card != nil {
		return card, nil
	}

	// 2. Запрашиваем у 1С: GET /medical-card/{patientID}
	path := "/medical-card/" + patientID
	body, err := u.onecClient.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	// 3. Парсим ответ
	var patientCard models.PatientCard
	if err := json.Unmarshal(body, &patientCard); err != nil {
		return nil, fmt.Errorf("failed to unmarshal medical card: %w", err)
	}

	// 4. Сохраняем в кеш
	_ = u.cacheRepo.SaveMedicalCard(ctx, patientID, &patientCard)

	return &patientCard, nil
}

func (u *MedCardUsecase) UpdateMedicalCard(ctx context.Context, req *models.UpdateMedicalCardRequest) error {
	// 1. Отправляем обновление в 1С
	path := "/medical-card/" + req.PatientID
	_, err := u.onecClient.Post(ctx, path, req)
	if err != nil {
		return fmt.Errorf("failed to update medical card in 1C: %w", err)
	}

	err = u.cacheRepo.DeleteMedicalCard(ctx, req.PatientID)
	if err != nil {
		// Логгируем, но не прерываем — главное, что 1С обновил
		// u.logger.Warnf("Failed to invalidate cache for patient %s: %v", req.PatientID, err)
	}

	return nil
}
