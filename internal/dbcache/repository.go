package dbcache

import (
	"context"
	"time"
)

type Repository interface {
	Read(ctx context.Context, key string) (string, error)
	Create(ctx context.Context, key, value string, expiration time.Time) error
}
