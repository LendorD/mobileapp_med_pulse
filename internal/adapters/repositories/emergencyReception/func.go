package emergencyReception

import (
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"

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

func (r *EmergencyReceptionRepositoryImpl) GetEmergencyReceptionsByDoctorAndDate(doctorID uint, date time.Time, page, perPage int) ([]models.EmergencyReceptionShortResponse, error) {
	var response []struct {
		Date        time.Time
		Status      string
		PatientName string
		Priority    bool
		Address     string
	}

	offset := (page - 1) * perPage
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.Model(&entities.EmergencyCall{}).
		Select(`
            emergency_receptions.date,
            emergency_receptions.status,
            emergency_receptions.priority,
            emergency_receptions.address,
            patients.full_name as patient_name
        `).
		Joins("LEFT JOIN patients ON patients.id = emergency_receptions.patient_id").
		Where("emergency_receptions.doctor_id = ? AND emergency_receptions.date >= ? AND emergency_receptions.date < ?",
			doctorID, startOfDay, endOfDay).
		Offset(offset).
		Limit(perPage).
		Order(getEmergencyOrderByPriorityAndDate()).
		Find(&response).
		Error

	// Преобразуем в финальную структуру
	result := make([]models.EmergencyReceptionShortResponse, len(response))
	for i, item := range response {
		result[i] = models.EmergencyReceptionShortResponse{
			Date:        item.Date.Format("2006-01-02 15:04"),
			Status:      item.Status,
			PatientName: item.PatientName,
			Priority:    item.Priority,
			Address:     item.Address,
		}
	}

	return result, err
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
