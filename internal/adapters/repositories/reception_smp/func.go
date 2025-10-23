package receptionSmp

import (
	"context"
	"encoding/json"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
)

func (r *ReceptionSmpRepositoryImpl) SaveReceptions(ctx context.Context, callID string, patients []models.Patient) error {
	_, err := json.Marshal(patients)
	if err != nil {
		return err
	}

	reception := entities.OneCReception{
		CallID: callID,
		Status: "received",
		// Receptions: data,
	}
	db := r.db.GetDB(ctx)
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Удаляем старую запись, если есть
		if err := tx.Where("call_id = ?", callID).Delete(&entities.OneCReception{}).Error; err != nil {
			return err
		}
		// Сохраняем новую
		return tx.Create(&reception).Error
	})
}

func (r *ReceptionSmpRepositoryImpl) GetReceptions(ctx context.Context, callID string) ([]models.Patient, error) {
	var reception entities.OneCReception
	db := r.db.GetDB(ctx)
	err := db.WithContext(ctx).Where("call_id = ?", callID).First(&reception).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var patients []models.Patient
	if err := json.Unmarshal(reception.MedServices, &patients); err != nil {
		return nil, err
	}
	return patients, nil
}
