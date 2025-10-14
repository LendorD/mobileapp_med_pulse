package interfaces

import "context"

type OneCClient interface {
	Get(ctx context.Context, path string) ([]byte, error)
	Post(ctx context.Context, path string, data interface{}) ([]byte, error)
}
