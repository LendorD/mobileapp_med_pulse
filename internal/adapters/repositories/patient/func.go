package patient

import (
	"context"
	"fmt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// СТАРЫЕ ФУНКЦИИ ------------------------------------------------------------------
func (r *PatientRepositoryImpl) GetPatientByIDWithTx(tx *gorm.DB, id uint) (*entities.Patient, error) {
	op := "repo.Patient.UpdatePatientWithTx"
	var patient entities.Patient
	if err := tx.First(&patient, id).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return &patient, nil
}

func (r *PatientRepositoryImpl) UpdatePatientWithTx(tx *gorm.DB, id uint, updateMap map[string]interface{}) (uint, error) {
	op := "repo.Patient.UpdatePatientWithTx"

	var updatedPatient entities.Patient
	result := tx.
		Clauses(clause.Returning{}).
		Model(&updatedPatient).
		Where("id = ?", id).
		Updates(updateMap)

	if result.Error != nil {
		return 0, errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, errors.NewNotFoundError("patient not found")
	}

	return updatedPatient.ID, nil
}

func (r *PatientRepositoryImpl) CreatePatient(patient entities.Patient) (uint, error) {
	op := "repo.Patient.CreatePatient"

	if err := r.db.Create(&patient).Error; err != nil {
		return 0, errors.NewDBError(op, fmt.Errorf("failed to create Patient: %w", err))
	}

	return patient.ID, nil
}

func (r *PatientRepositoryImpl) UpdatePatient(id uint, updateMap map[string]interface{}) (uint, error) {
	op := "repo.Patient.UpdatePatient"

	var updatedPatient entities.Patient
	result := r.db.
		Clauses(clause.Returning{}).
		Model(&updatedPatient).
		Where("id = ?", id).
		Updates(updateMap)

	if result.Error != nil {
		return 0, errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, errors.NewDBError(op, result.Error)
	}

	return updatedPatient.ID, nil
}

func (r *PatientRepositoryImpl) DeletePatient(id uint) error {
	op := "repo.Patient.ForceDeletePatient"

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Удаляем самые "глубокие" связи
	if err := tx.Exec("DELETE FROM reception_smp_med_services WHERE reception_smp_id IN (SELECT id FROM reception_smps WHERE patient_id = ?)", id).Error; err != nil {
		tx.Rollback()
		return errors.NewDBError(op, fmt.Errorf("failed to delete from reception_smp_med_services: %w", err))
	}

	// 2. Удаляем зависимости 1 уровня
	if err := tx.Exec("DELETE FROM reception_smps WHERE patient_id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.NewDBError(op, fmt.Errorf("failed to delete from reception_smps: %w", err))
	}

	if err := tx.Exec("DELETE FROM reception_hospitals WHERE patient_id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.NewDBError(op, fmt.Errorf("failed to delete from reception_hospitals: %w", err))
	}

	if err := tx.Exec("DELETE FROM patient_allergy WHERE patient_id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.NewDBError(op, fmt.Errorf("failed to delete from patient_allergy: %w", err))
	}

	// 3. Удаляем самого пациента
	result := tx.Delete(&entities.Patient{}, id)
	if result.Error != nil {
		tx.Rollback()
		return errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.NewDBError(op, errors.NewNotFoundError("patient not found"))
	}

	// 4. Коммитим транзакцию
	if err := tx.Commit().Error; err != nil {
		return errors.NewDBError(op, fmt.Errorf("failed to commit transaction: %w", err))
	}

	return nil
}

func (r *PatientRepositoryImpl) GetPatientByID(id uint) (entities.Patient, error) {
	op := "repo.Patient.GetPatientByID"

	var patient entities.Patient
	err := r.db.
		Preload("PersonalInfo").
		Preload("ContactInfo").
		Preload("Allergy").
		First(&patient, id).Error

	if err == gorm.ErrRecordNotFound {
		return entities.Patient{}, errors.NewNotFoundError("patient not found")
	}

	if err != nil {
		return entities.Patient{}, errors.NewDBError(op, err)
	}

	return patient, nil
}

func (r *PatientRepositoryImpl) GetPatientsByFullName(name string) ([]entities.Patient, error) {
	op := "repo.Patient.GetPatientsByFullName"

	var patients []entities.Patient
	if err := r.db.
		Where("full_name ILIKE ?", "%"+name+"%").
		Preload("PersonalInfo").
		Preload("ContactInfo").
		Find(&patients).Error; err != nil {

		return nil, errors.NewDBError(op, err)
	}

	if len(patients) == 0 {
		return nil, errors.NewDBError(op, errors.NewNotFoundError("no patients found"))
	}

	return patients, nil
}

func (r *PatientRepositoryImpl) GetPatientAllergiesByID(patientID uint) ([]entities.Allergy, error) {
	op := "repo.Allergy.GetPatientAllergiesByID"

	var patient entities.Patient
	err := r.db.
		Preload("Allergy").
		First(&patient, patientID).Error

	if err != nil {
		return nil, errors.NewDBError(op, err)
	}

	return patient.Allergy, nil
}

func (r *PatientRepositoryImpl) GetAllPatients(page, count int, queryFilter string, queryOrder string, parameters []interface{}) ([]entities.Patient, int64, error) {
	op := "repo.Patient.GetAllPatients"

	// Создаем базовый запрос
	query := r.db.Model(&entities.Patient{})

	// Применяем фильтрацию
	if queryFilter != "" {
		query = query.Where(queryFilter, parameters...)
	}

	// Применяем сортировку
	if queryOrder != "" {
		query = query.Order(queryOrder)
	}

	// Подсчитываем общее количество записей
	var totalRecords int64
	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, errors.NewDBError(op, err)
	}

	// Применяем пагинацию
	if page > 0 && count > 0 {
		offset := (page - 1) * count
		query = query.Offset(offset).Limit(count)
	}

	// Получаем записи
	var patients []entities.Patient
	result := query.Find(&patients)
	if result.Error != nil {
		return nil, 0, errors.NewDBError(op, result.Error)
	}

	return patients, totalRecords, nil
}
