package onecRepo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

const (
	patientListCacheKey  = "onec:patient_list"
	medicalCardKeyPrefix = "onec:medcard:patient:"
	cacheExpiry          = 24 * time.Hour
)

func (r *RedisOneCCacheRepository) SaveReceptions(ctx context.Context, callID string, patients []models.Patient) error {
	key := "onec:receptions:call:" + callID // безопасно для string
	data, err := json.Marshal(patients)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, cacheExpiry).Err()
}

func (r *RedisOneCCacheRepository) GetReceptions(ctx context.Context, callID string) ([]models.Patient, error) {
	key := "onec:receptions:call:" + callID
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var patients []models.Patient
	err = json.Unmarshal([]byte(val), &patients)
	return patients, err
}

// SavePatientList сохраняет список пациентов в Redis
func (r *RedisOneCCacheRepository) SavePatientList(ctx context.Context, patients []models.PatientListItem) error {
	// Удаляем старый список
	if err := r.client.Del(ctx, patientListCacheKey).Err(); err != nil {
		return err
	}

	// Преобразуем каждого пациента в JSON и добавляем в список
	var members []interface{}
	for _, p := range patients {
		data, err := json.Marshal(p)
		if err != nil {
			return err
		}
		members = append(members, data)
	}

	// Добавляем все элементы в список
	if len(members) > 0 {
		if err := r.client.RPush(ctx, patientListCacheKey, members...).Err(); err != nil {
			return err
		}
	}

	// Устанавливаем TTL при необходимости
	return r.client.Expire(ctx, patientListCacheKey, cacheExpiry).Err()
}

// GetPatientListPage возвращает список пациентов из Redis с пагинацией
func (r *RedisOneCCacheRepository) GetPatientListPage(ctx context.Context, offset, limit int) ([]models.PatientListItem, error) {
	end := offset + limit - 1
	values, err := r.client.LRange(ctx, patientListCacheKey, int64(offset), int64(end)).Result()
	if err != nil {
		return nil, err
	}

	var patients []models.PatientListItem
	for _, v := range values {
		var p models.PatientListItem
		if err := json.Unmarshal([]byte(v), &p); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}

func (r *RedisOneCCacheRepository) SaveMedicalCard(ctx context.Context, patientID string, card *models.PatientCard) error {
	key := medicalCardKeyPrefix + patientID
	data, err := json.Marshal(card)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, cacheExpiry).Err()
}

func (r *RedisOneCCacheRepository) GetMedicalCard(ctx context.Context, patientID string) (*models.PatientCard, error) {
	key := medicalCardKeyPrefix + patientID
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var card models.PatientCard
	if err := json.Unmarshal([]byte(val), &card); err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *RedisOneCCacheRepository) DeleteMedicalCard(ctx context.Context, patientID string) error {
	key := "onec:medcard:patient:" + patientID
	return r.client.Del(ctx, key).Err()
}
