package reception

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
)

func (r *ReceptionRepositoryImpl) CreateReception(reception *entities.Reception) error {
	return r.db.Create(reception).Error
}

func (r *ReceptionRepositoryImpl) UpdateReception(reception *entities.Reception) error {
	return r.db.Save(reception).Error
}

func (r *ReceptionRepositoryImpl) DeleteReception(id uint) error {
	return r.db.Delete(&entities.Reception{}, id).Error
}

func (r *ReceptionRepositoryImpl) GetReceptionByID(id uint) (*entities.Reception, error) {
	var reception entities.Reception
	if err := r.db.First(&reception, id).Error; err != nil {
		return nil, err
	}
	return &reception, nil
}

func (r *ReceptionRepositoryImpl) GetReceptionByDoctorID(doctorID uint) ([]entities.Reception, error) {
	var receptions []entities.Reception
	err := r.db.Where("doctor_id = ?", doctorID).Find(&receptions).Error
	return receptions, err
}

func (r *ReceptionRepositoryImpl) GetReceptionByPatientID(patientID uint) ([]entities.Reception, error) {
	var receptions []entities.Reception
	err := r.db.Where("patient_id = ?", patientID).Find(&receptions).Error
	return receptions, err
}

func (r *ReceptionRepositoryImpl) GetReceptionByDateRange(start, end time.Time) ([]entities.Reception, error) {
	var receptions []entities.Reception
	err := r.db.Where("date BETWEEN ? AND ?", start, end).Find(&receptions).Error
	return receptions, err
}

func getReceptionPriority(status entities.ReceptionStatus) int {
	switch status {
	case entities.StatusScheduled:
		return 1
	case entities.StatusCompleted:
		return 2
	case entities.StatusCancelled, entities.StatusNoShow:
		return 3
	default:
		return 4
	}
}

func getOrderByStatusAndDate() string {
	return `
        CASE 
            WHEN status = 'emergency' THEN 1
            WHEN status = 'scheduled' THEN 2
            WHEN status = 'completed' THEN 3
            WHEN status = 'cancelled' THEN 4
            ELSE 5
        END,
        date ASC
    `
}

func (r *ReceptionRepositoryImpl) GetReceptionsByDoctorAndDate(doctorID uint, date time.Time, page, perPage int) ([]models.ReceptionShortResponse, error) {
	var response []struct {
		Date        time.Time
		Status      string
		PatientName string
		IsOut       bool
	}

	offset := (page - 1) * perPage
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.Model(&entities.Reception{}).
		Select(`
            receptions.date,
            receptions.status,
            receptions.is_out,
            patients.full_name as patient_name
        `).
		Joins("LEFT JOIN patients ON patients.id = receptions.patient_id").
		Where("receptions.doctor_id = ? AND receptions.date >= ? AND receptions.date < ?",
			doctorID, startOfDay, endOfDay).
		Offset(offset).
		Limit(perPage).
		Order(getOrderByStatusAndDate()).
		Find(&response).
		Error

	// Преобразуем в финальную структуру с форматированной датой
	result := make([]models.ReceptionShortResponse, len(response))
	for i, item := range response {
		result[i] = models.ReceptionShortResponse{
			Date:        item.Date.Format("2006-01-02 15:04"),
			Status:      item.Status,
			PatientName: item.PatientName,
			IsOut:       item.IsOut,
		}
	}

	return result, err
}
