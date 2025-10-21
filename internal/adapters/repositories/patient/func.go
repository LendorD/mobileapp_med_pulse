package patient

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SavePatientList сохраняет полный список пациентов (заменяет текущий)
func (r *PatientRepositoryImpl) SavePatientList(ctx context.Context, patients []entities.OneCPatientListItem) error {
	db := r.db.GetDB(ctx)
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

// SaveOrUpdatePatientList сохраняет список пациентов, обновляя существующих и добавляя новых
func (r *PatientRepositoryImpl) SaveOrUpdatePatientList(ctx context.Context, patients []entities.OneCPatientListItem) error {
	if len(patients) == 0 {
		return nil
	}
	db := r.db.GetDB(ctx)
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Вставляем или обновляем существующих по PatientID
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "patient_id"}}, // уникальное поле
			UpdateAll: true,                                  // обновляем все поля при конфликте
		}).CreateInBatches(patients, 100).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetPatientListPage возвращает страницу пациентов
func (r *PatientRepositoryImpl) GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, int64, error) {
	var patients []entities.OneCPatientListItem
	var total int64
	db := r.db.GetDB(ctx)
	// Получаем общее количество
	if err := db.WithContext(ctx).Model(&entities.OneCPatientListItem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Получаем страницу
	if err := db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}
