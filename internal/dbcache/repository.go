package dbcache

import (
	"context"
	"time"
)

type Repository interface {
	Read(ctx context.Context, key string) (string, time.Time, error)
	Create(ctx context.Context, key, value string, expiration time.Time) error
	Update(ctx context.Context, key, value string, expiration time.Time) error
}
