package onecRepo

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/redis/go-redis/v9"
)

type RedisOneCCacheRepository struct {
	client *redis.Client
}

func NewRedisOneCCacheRepository(client *redis.Client) interfaces.OneCCacheRepository {
	return &RedisOneCCacheRepository{client: client}
}
