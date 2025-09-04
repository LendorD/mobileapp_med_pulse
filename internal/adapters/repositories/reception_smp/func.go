package receptionSmp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *ReceptionSmpRepositoryImpl) CreateReceptionSmp(reception entities.ReceptionSMP) (uint, error) {
	op := "repo.ReceptionSmp.CreateReceptionSmp"
	if err := r.db.Clauses(clause.Returning{}).Create(&reception).Error; err != nil {
		return 0, errors.NewDBError(op, err)
	}
	return reception.ID, nil
}

func (r *ReceptionSmpRepositoryImpl) UpdateReceptionSmp(id uint, updateMap map[string]interface{}) (uint, error) {
	op := "repo.ReceptionSmp.UpdateReceptionSmp"

	var updatedReception entities.ReceptionSMP
	tx := r.db.Begin()

	// Обработка связи med_services отдельно
	var medServiceIDs []uint
	if raw, ok := updateMap["med_services"]; ok {
		if arr, ok := raw.([]interface{}); ok {
			for _, id := range arr {
				if f64, ok := id.(float64); ok {
					medServiceIDs = append(medServiceIDs, uint(f64))
				}
			}
		}
		delete(updateMap, "med_services") // удаляем, чтобы не было попытки обновить колонку
	}

	// Обновление обычных полей
	result := tx.
		Clauses(clause.Returning{}).
		Model(&updatedReception).
		Where("id = ?", id).
		Updates(updateMap)

	if result.Error != nil {
		tx.Rollback()
		return 0, errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return 0, errors.NewNotFoundError("reception not found")
	}

	// Обновление связей many2many
	if len(medServiceIDs) > 0 {
		var medServices []entities.MedService
		if err := tx.Where("id IN ?", medServiceIDs).Find(&medServices).Error; err != nil {
			tx.Rollback()
			return 0, errors.NewDBError(op+": find med_services", err)
		}
		if err := tx.Model(&updatedReception).Association("MedServices").Replace(&medServices); err != nil {
			tx.Rollback()
			return 0, errors.NewDBError(op+": replace med_services", err)
		}
	}

	returnedID := updatedReception.ID
	return returnedID, tx.Commit().Error
}

func (r *ReceptionSmpRepositoryImpl) DeleteReceptionSmp(id uint) error {
	op := "repo.ReceptionSmp.DeleteReceptionSmp"
	result := r.db.Delete(&entities.ReceptionSMP{}, id)
	if result.Error != nil {
		return errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("reception not found")
	}
	return nil
}

func (r *ReceptionSmpRepositoryImpl) GetReceptionSmpByID(id uint) (entities.ReceptionSMP, error) {
	op := "repo.ReceptionSmp.GetReceptionSmpByID"

	var reception entities.ReceptionSMP
	if err := r.db.Preload("MedServices").First(&reception, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ReceptionSMP{}, errors.NewNotFoundError("reception not found")
		}
		return entities.ReceptionSMP{}, errors.NewDBError(op, err)
	}
	return reception, nil
}

func (r *ReceptionSmpRepositoryImpl) GetReceptionSmpByDoctorID(doctorID uint) ([]entities.ReceptionSMP, error) {
	op := "repo.ReceptionSmp.GetReceptionSmpByDoctorID"

	var receptions []entities.ReceptionSMP
	if err := r.db.Where("doctor_id = ?", doctorID).Find(&receptions).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return receptions, nil
}

func (r *ReceptionSmpRepositoryImpl) GetReceptionSmpByPatientID(
	patientID uint,
	page, count int,
	filter, order string,
	params []interface{},
) ([]entities.ReceptionSMP, int64, error) {
	op := "repo.ReceptionSmp.GetReceptionSmpByPatientID"

	var (
		receptions []entities.ReceptionSMP
		totalRows  int64
	)

	// Базовый запрос
	query := r.db.Model(&entities.ReceptionSMP{}).Where("patient_id = ?", patientID)

	// Применяем фильтр, если есть
	if filter != "" {
		query = query.Where(filter, params...)
	}

	// Считаем общее количество строк (до пагинации)
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, 0, errors.NewDBError(op, err)
	}

	// Применяем сортировку
	if order != "" {
		query = query.Order(order)
	}

	// Применяем пагинацию (если включена)
	if count > 0 {
		offset := (page - 1) * count
		query = query.Offset(offset).Limit(count)
	}

	// Получаем данные
	if err := query.Preload("Doctor").Preload("Patient").Find(&receptions).Error; err != nil {
		return nil, 0, errors.NewDBError(op, err)
	}

	return receptions, totalRows, nil
}

func (r *ReceptionSmpRepositoryImpl) GetReceptionSmpByDateRange(start, end time.Time) ([]entities.ReceptionSMP, error) {
	op := "repo.ReceptionSmp.GetReceptionSmpByDateRange"

	var receptions []entities.ReceptionSMP
	if err := r.db.Where("date BETWEEN ? AND ?", start, end).Find(&receptions).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return receptions, nil
}

func (r *ReceptionSmpRepositoryImpl) UpdateReceptionSmpMedServices(receptionID uint, services []entities.MedService) error {
	op := "repo.ReceptionSmp.UpdateReceptionSmpMedServices"
	if len(services) == 0 {
		return nil
	}
	// Получаем ID всех услуг для проверки их существования
	serviceIDs := make([]uint, len(services))
	for i, s := range services {
		serviceIDs[i] = s.ID
	}

	// Проверяем что все услуги существуют
	var count int64
	if err := r.db.Model(&entities.MedService{}).Where("id IN ?", serviceIDs).Count(&count).Error; err != nil {
		return errors.NewDBError(op, err)
	}
	if int(count) != len(serviceIDs) {
		return errors.NewDBError(op, fmt.Errorf("some med services not found"))
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Удаляем старые связи
	if err := tx.Exec("DELETE FROM reception_smp_med_services WHERE reception_smp_id = ?", receptionID).Error; err != nil {
		tx.Rollback()
		return errors.NewDBError(op, err)
	}

	// Создаём batch для вставки
	var inserts []map[string]interface{}
	for _, id := range serviceIDs {
		inserts = append(inserts, map[string]interface{}{
			"reception_smp_id": receptionID,
			"med_service_id":   id,
		})
	}

	// Вставляем новые связи batch-ом
	if err := tx.Table("reception_smp_med_services").Create(inserts).Error; err != nil {
		tx.Rollback()
		return errors.NewDBError(op, err)
	}

	if err := tx.Commit().Error; err != nil {
		return errors.NewDBError(op, err)
	}

	return nil
}

func (r *ReceptionSmpRepositoryImpl) SaveSignature(receptionID uint, signature []byte) error {
	return r.db.Model(&entities.ReceptionSMP{}).
		Where("id = ?", receptionID).
		Update("patient_signature", signature).
		Error
}

// Получаем подпись для конкретного медицинского заключения
func (r *ReceptionSmpRepositoryImpl) GetSignature(receptionID uint) ([]byte, error) {
	var reception entities.ReceptionSMP
	if err := r.db.Select("patient_signature").First(&reception, receptionID).Error; err != nil {
		return nil, err
	}
	return reception.PatientSignature, nil
}

// Обновленные методы репозитория
func (r *ReceptionSmpRepositoryImpl) GetWithPatientsByEmergencyCallID(
	emergencyCallID uint,
	page, perPage int,
) ([]entities.ReceptionSMP, int64, error) {
	op := "repo.ReceptionSmp.GetReceptionSmpByPatientID"

	var receptions []entities.ReceptionSMP
	var total int64

	baseQuery := r.db.Model(&entities.ReceptionSMP{}).
		Where("emergency_call_id = ?", emergencyCallID)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, errors.NewDBError(op, err)
	}

	offset := (page - 1) * perPage
	err := baseQuery.
		Preload("Patient").
		Preload("MedServices").
		Preload("Doctor").
		Order("created_at DESC").
		Offset(offset).
		Limit(perPage).
		Find(&receptions).
		Error

	if err != nil {
		return nil, 0, errors.NewDBError(op, err)
	}

	for i := range receptions {

		receptions[i].SpecializationDataDecoded = pgtype.JSONB{
			Bytes:  []byte(`{"key":"value"}`),
			Status: pgtype.Present,
		}
	}

	return receptions, total, nil
}
func (r *ReceptionSmpRepositoryImpl) GetReceptionWithMedServicesByID(
	smpID uint,
	callID uint,
) (entities.ReceptionSMP, error) {
	op := "repo.ReceptionSmp.GetReceptionWithMedServicesByID"
	var reception entities.ReceptionSMP

	// Явно указываем нужные поля для загрузки
	query := r.db.
		Preload("Patient").
		Preload("MedServices").
		Preload("Doctor.Specialization").
		Where("id = ?", smpID)

	if callID > 0 {
		query = query.Where("emergency_call_id = ?", callID)
	}

	err := query.First(&reception).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ReceptionSMP{}, errors.NewNotFoundError("reception not found")
		}
		return entities.ReceptionSMP{}, errors.NewDBError(op, err)
	}

	// Улучшенное декодирование JSONB
	if err := decodeReceptionSpecializationData(&reception); err != nil {
		log.Printf("%s: failed to decode specialization data: %v", op, err)
		// Не возвращаем ошибку, так как основная информация уже получена
	}

	return reception, nil
}

// Вынесенная функция для декодирования
func decodeReceptionSpecializationData(reception *entities.ReceptionSMP) error {
	// Проверяем что данные существуют и не пустые
	if reception.SpecializationData.Status != pgtype.Present ||
		len(reception.SpecializationData.Bytes) == 0 {
		return nil
	}

	// Получаем название специализации (с проверкой)
	specTitle := ""
	if reception.Doctor.Specialization.ID != 0 { // Проверка что специализация загружена
		specTitle = reception.Doctor.Specialization.Title
	}

	decoded, err := decodeSpecializationData(reception.SpecializationData, specTitle)
	if err != nil {
		return fmt.Errorf("decoding failed: %w", err)
	}

	reception.SpecializationDataDecoded = decoded
	return nil
}

// Обновлённая функция декодирования
func decodeSpecializationData(data pgtype.JSONB, specialization string) (interface{}, error) {
	// Проверка наличия данных
	if data.Status != pgtype.Present || len(data.Bytes) == 0 {
		fmt.Print("BEDAAAA")
		return nil, nil
	}

	var result interface{}
	switch specialization {
	case "Невролог":
		result = new(entities.NeurologistData)
	case "Травматолог":
		result = new(entities.TraumatologistData)
	case "Психиатр":
		result = new(entities.PsychiatristData)
	case "Уролог":
		result = new(entities.UrologistData)
	case "Оториноларинголог":
		result = new(entities.OtolaryngologistData)
	case "Проктолог":
		result = new(entities.ProctologistData)
	case "Аллерголог":
		result = new(entities.AllergologistData)
	default:
		result = make(map[string]interface{})
	}

	// Декодирование с проверкой структуры
	decoder := json.NewDecoder(bytes.NewReader(data.Bytes))
	decoder.DisallowUnknownFields() // Для отлова несоответствий структур

	if err := decoder.Decode(&result); err != nil {
		log.Printf("Decoding error for %s: %v, data: %s", specialization, err, string(data.Bytes))
		return nil, fmt.Errorf("decoding error: %w", err)
	}

	return result, nil
}
