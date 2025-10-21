package usecases

import (
	"context"
	"fmt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

type MedCardUsecase struct {
	repo       interfaces.MedicalCardRepository
	onecClient interfaces.OneCClient
	txManager  interfaces.TxManager
}

func NewMedCardUsecase(
	repo interfaces.MedicalCardRepository,
	onecClient interfaces.OneCClient,
	txManager interfaces.TxManager,
) interfaces.MedCardUsecase {
	return &MedCardUsecase{
		repo:       repo,
		onecClient: onecClient,
		txManager:  txManager,
	}
}

// GetMedCardByPatientID — получает карту пациента из БД или из 1С
func (u *MedCardUsecase) GetMedCardByPatientID(ctx context.Context, patientID string) (*entities.OneCMedicalCard, error) {
	card, err := u.repo.GetMedicalCard(ctx, patientID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("db error: %w", err)
	}
	if card != nil {
		return card, nil
	}

	OneCCard, err := u.onecClient.GetMedCardByPatientID(patientID)
	if err := u.repo.SaveMedicalCard(ctx, OneCCard); err != nil {
		fmt.Printf("warn: failed to save medical card for patient %s: %v\n", patientID, err)
		return nil, fmt.Errorf("failed to save medical card %v", err)
	}

	return OneCCard, nil
}

// UpdateMedicalCard — обновляет карту в 1С и БД
func (u *MedCardUsecase) UpdateMedicalCard(ctx context.Context, card *entities.OneCMedicalCard) error {
	if err := u.onecClient.UpdateMedCardByPatientID(card.PatientID, card); err != nil {
		return fmt.Errorf("failed to update in 1C: %w", err)
	}

	if err := u.repo.SaveMedicalCard(ctx, card); err != nil {
		fmt.Printf("warn: failed to update local cache for patient %s: %v\n", card.PatientID, err)
	}

	return nil
}
