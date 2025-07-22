package allergy

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *AllergyRepositoryImpl) CreateAllergy(allergy *entities.Allergy) (uint, error) {
	op := "repo.Allergy.CreateAllergy"

	err := r.db.Clauses(clause.Returning{}).Create(&allergy).Error
	if err != nil {
		return 0, errors.NewDBError(op, err)
	}
	return allergy.ID, nil
}

func (r *AllergyRepositoryImpl) UpdateAllergy(id uint, updateMap map[string]interface{}) (uint, error) {
	op := "repo.Allergy.UpdateAllergy"

	var updatedAllergy entities.Allergy
	result := r.db.
		Clauses(clause.Returning{}).
		Model(&updatedAllergy).
		Where("id = ?", id).
		Updates(updateMap)

	if result.Error != nil {
		return 0, errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, errors.NewDBError(op, result.Error)
	}

	return updatedAllergy.ID, nil
}

func (r *AllergyRepositoryImpl) DeleteAllergy(id uint) error {
	op := "repo.Allergy.DeleteAllergy"

	result := r.db.Delete(&entities.Allergy{}, id)
	if result.Error != nil {
		return errors.NewDBError(op, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewDBError(op, result.Error)
	}

	return nil
}

func (r *AllergyRepositoryImpl) GetAllergyByID(id uint) (entities.Allergy, error) {
	op := "repo.Allergy.GetAllergyByID"

	var allergy entities.Allergy
	err := r.db.
		First(&allergy, id).Error
	if err != nil {
		return entities.Allergy{}, errors.NewDBError(op, err)
	}

	return allergy, nil
}

func (r *AllergyRepositoryImpl) GetAllergyByName(name string) (entities.Allergy, error) {
	op := "repo.Allergy.GetAllergyByName"

	var allergy entities.Allergy
	err := r.db.
		Where("name = ?", name).
		First(&allergy).Error

	if err != nil {
		return entities.Allergy{}, errors.NewDBError(op, err)
	}

	return allergy, nil
}

func (r *AllergyRepositoryImpl) GetAllAllergies() ([]entities.Allergy, error) {
	op := "repo.Allergy.GetAllAllergies"

	var allergies []entities.Allergy
	if err := r.db.Find(&allergies).Error; err != nil {
		return nil, errors.NewDBError(op, err)
	}

	return allergies, nil
}

func (r *AllergyRepositoryImpl) GetAllergiesByPatientID(patientID uint) ([]entities.Allergy, error) {
	op := "repo.Allergy.GetAllAllergiesByPatientId"
	var allergies []entities.Allergy
	err := r.db.Model(&entities.Patient{ID: patientID}).
		Association("Allergy").
		Find(&allergies)
	if err == nil {
		return nil, errors.NewDBError(op, err)
	}
	return allergies, err
}

func (r *AllergyRepositoryImpl) RemovePatientAllergies(patientID uint, allergies []entities.Allergy) error {
	return r.db.Model(&entities.Patient{ID: patientID}).
		Association("Allergy").
		Delete(allergies)
}

func (r *AllergyRepositoryImpl) AddPatientAllergies(patientID uint, allergies []entities.Allergy) error {
	return r.db.Model(&entities.Patient{ID: patientID}).
		Association("Allergy").
		Append(allergies)
}

func (r *AllergyRepositoryImpl) GetPatientAllergiesByIDWithTx(tx *gorm.DB, patientID uint) ([]entities.Allergy, error) {
	var allergies []entities.Allergy
	if err := tx.Model(&entities.Patient{ID: patientID}).
		Association("Allergy").
		Find(&allergies); err != nil {
		return nil, err
	}
	return allergies, nil
}

func (r *AllergyRepositoryImpl) SyncPatientAllergiesWithTx(tx *gorm.DB, patientID uint, allergies []entities.Allergy) error {
	op := "repo.Allergy.SyncPatientAllergiesWithTx"

	// Получаем только названия текущих аллергий
	var currentNames []string
	if err := tx.Model(&entities.Allergy{}).
		Joins("JOIN patient_allergy ON patient_allergy.allergy_id = allergies.id").
		Where("patient_allergy.patient_id = ?", patientID).
		Pluck("name", &currentNames).Error; err != nil {
		return errors.NewDBError(op, err)
	}

	// Создаем мапы для сравнения
	currentMap := make(map[string]bool)
	for _, name := range currentNames {
		currentMap[name] = true
	}

	newMap := make(map[string]bool)
	for _, a := range allergies {
		newMap[a.Name] = true
	}

	// Удаляем отсутствующие в новом списке
	var toRemove []string
	for name := range currentMap {
		if !newMap[name] {
			toRemove = append(toRemove, name)
		}
	}

	if len(toRemove) > 0 {
		if err := tx.Exec(`
            DELETE FROM patient_allergy 
            WHERE patient_id = ? 
            AND allergy_id IN (
                SELECT id FROM allergies WHERE name IN ?
            )`, patientID, toRemove).Error; err != nil {
			return errors.NewDBError(op, err)
		}
	}

	// Добавляем новые аллергии
	for _, allergy := range allergies {
		if currentMap[allergy.Name] {
			continue // Уже существует, пропускаем
		}

		// Проверяем существование аллергии в базе
		var existingAllergy entities.Allergy
		if err := tx.Where("name = ?", allergy.Name).First(&existingAllergy).Error; err == nil {
			// Используем существующую аллергию
			allergy = existingAllergy
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// Создаем новую аллергию
			if err := tx.Create(&allergy).Error; err != nil {
				return errors.NewDBError(op, err)
			}
		} else {
			return errors.NewDBError(op, err)
		}

		// Добавляем связь
		if err := tx.Exec(`
            INSERT INTO patient_allergy (patient_id, allergy_id) 
            VALUES (?, ?)
            ON CONFLICT (patient_id, allergy_id) DO NOTHING`,
			patientID, allergy.ID).Error; err != nil {
			return errors.NewDBError(op, err)
		}
	}

	return nil
}
