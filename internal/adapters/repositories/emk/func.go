package emk

import (
	"context"
	"errors"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"gorm.io/gorm"
)

// SaveEmkList — сохраняет список ЭМК для конкретного пациента
func (r *EmkRepository) SaveEmkList(ctx context.Context, patientID string, emkList []entities.Emk) error {
	db := r.db.GetDB(ctx)
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Удаляем старые записи по patient_id
		if err := tx.Where("patient_id = ?", patientID).Delete(&entities.Emk{}).Error; err != nil {
			return err
		}
		// Вставляем новые записи
		if len(emkList) > 0 {
			return tx.CreateInBatches(emkList, 100).Error
		}
		return nil
	})
}

// GetEmkByPatientID — получает список ЭМК по patient_id
func (r *EmkRepository) GetEmkByPatientID(ctx context.Context, patientID string) ([]entities.Emk, error) {
	var emkList []entities.Emk
	db := r.db.GetDB(ctx)
	err := db.WithContext(ctx).
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&emkList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return emkList, err
}
