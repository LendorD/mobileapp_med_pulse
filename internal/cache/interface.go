package cache

import "context"

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl int) error // ttl в секундах
	Del(ctx context.Context, key string) error
}
