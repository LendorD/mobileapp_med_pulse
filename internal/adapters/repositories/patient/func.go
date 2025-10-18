package patient

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"gorm.io/gorm"
)

// SavePatientList сохраняет полный список пациентов (заменяет текущий)
func (r *PatientRepositoryImpl) SavePatientList(ctx context.Context, patients []entities.OneCPatientListItem) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Очищаем таблицу
		if err := tx.Delete(&entities.OneCPatientListItem{}, "1=1").Error; err != nil {
			return err
		}
		// Вставляем новые данные
		if len(patients) > 0 {
			return tx.CreateInBatches(patients, 100).Error
		}
		return nil
	})
}

// GetPatientListPage возвращает страницу пациентов
func (r *PatientRepositoryImpl) GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, int64, error) {
	var patients []entities.OneCPatientListItem
	var total int64

	// Получаем общее количество
	if err := r.db.WithContext(ctx).Model(&entities.OneCPatientListItem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Получаем страницу
	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}
