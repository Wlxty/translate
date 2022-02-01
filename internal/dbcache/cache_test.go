package dbcache

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"translateapp/internal/logger"
)

func TestConnection(t *testing.T) {
	_, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/postgres")
	assert.NotEqualf(t, err, nil, "Connection refused")

}
func TestGetCache(t *testing.T) {
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/postgres")
	logger := logger.NewLogger("debug", true)

	var rt = NewThroughDB(conn, logger)
	duration := time.Hour * 2
	value, _ := rt.Get("sample", Sample, duration)
	assert.Equalf(t, value, "Sample", "They should be equal")

}

func Sample() (interface{}, error) {
	return "Sample", nil
}
