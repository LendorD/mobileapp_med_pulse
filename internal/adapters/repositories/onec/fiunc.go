package onecRepo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

func (r *RedisOneCCacheRepository) SaveReceptions(ctx context.Context, callID int, receptions []models.Reception) error {
	key := fmt.Sprintf("onec:receptions:call:%d", callID)
	data, err := json.Marshal(receptions)
	if err != nil {
		return err
	}

	// Храним 1 час (или сколько нужно)
	return r.client.Set(ctx, key, data, 1*time.Hour).Err()
}

func (r *RedisOneCCacheRepository) GetReceptions(ctx context.Context, callID int) ([]models.Reception, error) {
	key := fmt.Sprintf("onec:receptions:call:%d", callID)
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // не найдено
	}
	if err != nil {
		return nil, err
	}

	var receptions []models.Reception
	err = json.Unmarshal([]byte(val), &receptions)
	return receptions, err
}
