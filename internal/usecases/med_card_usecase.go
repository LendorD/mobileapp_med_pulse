package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

type MedCardUsecase struct {
	repo       interfaces.MedicalCardRepository
	onecClient interfaces.OneCClient
}

func NewMedCardUsecase(
	repo interfaces.MedicalCardRepository,
	onecClient interfaces.OneCClient,
) interfaces.MedCardUsecase {
	return &MedCardUsecase{
		onecClient: onecClient,
	}
}

// GetMedCardByPatientID — получает карту из БД или 1С
func (u *MedCardUsecase) GetMedCardByPatientID(ctx context.Context, patientID string) (*entities.OneCMedicalCard, error) {
	// 1. Пробуем из БД
	card, err := u.repo.GetMedicalCard(ctx, patientID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("db error: %w", err)
	}
	if card != nil {
		return card, nil
	}

	// 2. Запрашиваем у 1С
	path := "/medical-card/" + patientID
	body, err := u.onecClient.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("1C error: %w", err)
	}

	// 3. Парсим в entities.OneCMedicalCard
	var patientCard entities.OneCMedicalCard
	if err := json.Unmarshal(body, &patientCard); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	// 4. Сохраняем в БД
	if err := u.repo.SaveMedicalCard(ctx, &patientCard); err != nil {
		// Логируем, но не прерываем
		fmt.Printf("warn: failed to cache medical card: %v", err)
	}

	return &patientCard, nil
}

// UpdateMedicalCard — обновляет карту в 1С и БД
func (u *MedCardUsecase) UpdateMedicalCard(ctx context.Context, card *entities.OneCMedicalCard) error {
	// 1. Отправляем в 1С
	path := "/medical-card/" + card.PatientID
	body, err := json.Marshal(card)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	_, err = u.onecClient.Post(ctx, path, body)
	if err != nil {
		return fmt.Errorf("1C update error: %w", err)
	}

	// 2. Сохраняем в БД (перезаписываем)
	if err := u.repo.SaveMedicalCard(ctx, card); err != nil {
		fmt.Printf("warn: failed to update cache: %v", err)
	}

	return nil
}
