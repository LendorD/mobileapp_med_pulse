package emergencyReception

import (
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
)

func (r *EmergencyReceptionRepositoryImpl) CreateEmergencyReception(er entities.EmergencyCall) error {
	op := "repo.EmergencyReception.CreateEmergencyReception"

	if err := r.db.Create(&er).Error; err != nil {
		return errors.NewDBError(op, err)
	}
	return nil
}

func (r *EmergencyReceptionRepositoryImpl) UpdateEmergencyReception(id uint, updateMap map[string]interface{}) (uint, error) {
	op := "repo.EmergencyReception.UpdateEmergencyReception"

	var updatedER entities.EmergencyCall
	result := r.db.
		Clauses(clause.Returning{}).
		Model(&updatedER).
		Where("id = ?", id).
		Updates(updateMap)

	if result.Error != nil {
		return 0, errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, errors.NewNotFoundError("emergency reception not found")
	}

	return updatedER.ID, nil
}

func (r *EmergencyReceptionRepositoryImpl) DeleteEmergencyReception(id uint) error {
	op := "repo.EmergencyReception.DeleteEmergencyReception"

	result := r.db.Delete(&entities.EmergencyCall{}, id)
	if result.Error != nil {
		return errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("emergency reception not found")
	}
	return nil
}
func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionByID(id uint) (entities.EmergencyCall, error) {
	op := "repo.EmergencyReception.GetEmergencyReceptionByID"

	var er entities.EmergencyCall
	if err := r.db.First(&er, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.EmergencyCall{}, errors.NewNotFoundError("emergency reception not found")
		}
		return entities.EmergencyCall{}, errors.NewDBError(op, err)
	}
	return entities.EmergencyCall{}, nil
}

func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionsByDoctorID(doctorID uint) ([]entities.EmergencyCall, error) {
	op := "repo.EmergencyReception.GetEmergencyReceptionByDoctorID"

	var list []entities.EmergencyCall
	if err := r.db.Where("doctor_id = ?", doctorID).Find(&list).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return list, nil
}

func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionsByPatientID(patientID uint) ([]entities.EmergencyCall, error) {
	op := "repo.EmergencyReception.GetEmergencyReceptionByPatientID"

	var list []entities.EmergencyCall
	if err := r.db.Where("patient_id = ?", patientID).Find(&list).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return list, nil
}

func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionsByDateRange(start, end time.Time) ([]entities.EmergencyCall, error) {
	op := "repo.EmergencyReception.GetEmergencyReceptionByDateRange"

	var list []entities.EmergencyCall
	if err := r.db.Where("date BETWEEN ? AND ?", start, end).Find(&list).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return list, nil
}

func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionsPriorityCases() ([]entities.EmergencyCall, error) {
	op := "repo.EmergencyReception.GetEmergencyReceptionPriorityCases"

	var list []entities.EmergencyCall
	if err := r.db.Where("priority = true").Find(&list).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}
	return list, nil
}

func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionsByDoctorAndDate(doctorID uint, date time.Time, page, perPage int) ([]models.EmergencyCallShortResponse, error) {
	var response []struct {
		CreatedAt time.Time `gorm:"column:created_at"`
		Status    string
		Phone     string
		Priority  bool
		Address   string
	}

	offset := (page - 1) * perPage

	query := r.db.Model(&entities.EmergencyCall{}).
		Select(`
            emergency_calls.created_at,
            emergency_calls.status,
            emergency_calls.phone,
            emergency_calls.priority,
            emergency_calls.address
        `).
		Where("emergency_calls.doctor_id = ?", doctorID)

	if !date.IsZero() {
		startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
		endOfDay := startOfDay.Add(24 * time.Hour)
		query = query.Where("DATE(created_at) BETWEEN ? AND ?",
			startOfDay.Format("2006-01-02"),
			endOfDay.Format("2006-01-02"))
	}

	err := query.
		Offset(offset).
		Limit(perPage).
		Order("priority DESC, created_at DESC"). // Простая сортировка по приоритету и дате
		Find(&response).
		Error

	if err != nil {
		return nil, fmt.Errorf("failed to get emergency receptions: %v", err)
	}

	result := make([]models.EmergencyCallShortResponse, len(response))
	for i, item := range response {
		result[i] = models.EmergencyCallShortResponse{
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
			Status:    item.Status,
			Phone:     item.Phone,
			Priority:  item.Priority,
			Address:   item.Address,
		}
	}

	return result, nil
}

// Вспомогательная функция для сортировки
func getEmergencyOrderByPriorityAndDate() string {
	return `
        CASE 
            WHEN priority = true THEN 1
            ELSE 2
        END,
        date ASC
    `
}
