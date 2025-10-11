package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &RedisCache{Client: rdb}
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // ключ не найден — не ошибка
	}
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value []byte, ttlSec int) error {
	ttl := time.Duration(ttlSec) * time.Second
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCache) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
