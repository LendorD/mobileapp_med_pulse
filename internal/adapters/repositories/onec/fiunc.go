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
	data, err := json.Marshal(patients)
	if err != nil {
		return err
	}

	// Сохраняем с TTL (можно убрать TTL, передав 0)
	return r.client.Set(ctx, patientListCacheKey, data, cacheExpiry).Err()
}

// GetPatientList возвращает список пациентов из Redis
func (r *RedisOneCCacheRepository) GetPatientList(ctx context.Context) ([]models.PatientListItem, error) {
	data, err := r.client.Get(ctx, patientListCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			// Ключ не найден — возвращаем пустой список
			return []models.PatientListItem{}, nil
		}
		return nil, err
	}

	var patients []models.PatientListItem
	if err := json.Unmarshal([]byte(data), &patients); err != nil {
		return nil, err
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
