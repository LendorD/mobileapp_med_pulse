package medcard

import (
	"context"
	"errors"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"gorm.io/gorm"
)

func (r *MedicalCardRepository) SaveMedicalCard(ctx context.Context, card *entities.OneCMedicalCard) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("patient_id = ?", card.PatientID).Delete(&entities.OneCMedicalCard{}).Error; err != nil {
			return err
		}
		return tx.Create(card).Error
	})
}

func (r *MedicalCardRepository) GetMedicalCard(ctx context.Context, patientID string) (*entities.OneCMedicalCard, error) {
	var card entities.OneCMedicalCard
	err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).First(&card).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &card, err
}

func (r *MedicalCardRepository) DeleteMedicalCard(ctx context.Context, patientID string) error {
	return r.db.WithContext(ctx).Where("patient_id = ?", patientID).Delete(&entities.OneCMedicalCard{}).Error
}
