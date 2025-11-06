package flg

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

// CreateFlg создаёт новую запись флюрографии
func (r *FlgRepository) CreateFlg(ctx context.Context, flg *entities.Flg) error {
	op := "repo.Flg.CreateFlg"
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Create(flg).Error; err != nil {
		return errors.NewDBError(op, err)
	}
	return nil
}

// GetFlgByPatientID возвращает все флюрографии пациента
func (r *FlgRepository) GetFlgByPatientID(ctx context.Context, patientID uint) ([]entities.Flg, error) {
	op := "repo.Flg.GetFlgByPatientID"
	var flgs []entities.Flg
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Where("patient_id = ?", patientID).Find(&flgs).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return flgs, nil
}

// GetFlgByID возвращает флюрографию по ID (может понадобиться позже)
func (r *FlgRepository) GetFlgByID(ctx context.Context, id uint) (*entities.Flg, error) {
	op := "repo.Flg.GetFlgByID"
	var flg entities.Flg
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Where("id = ?", id).First(&flg).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFoundError(op)
		}
		return nil, errors.NewDBError(op, err)
	}
	return &flg, nil
}

// Delete удаляет флюрографию по ID
func (r *FlgRepository) Delete(ctx context.Context, id uint) error {
	op := "repo.Flg.Delete"
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Delete(&entities.Flg{}, id).Error; err != nil {
		return errors.NewDBError(op, err)
	}
	return nil
}
