package dbcache

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/postgres")
	repo := NewRepo(conn)
	value, _ := repo.Read(context.Background(), "Sample")
	assert.Equalf(t, value, "probka", "They should be equal")
}

func TestCreate(t *testing.T) {
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/postgres")
	repo := NewRepo(conn)
	duration := time.Hour * 2
	expiration := time.Now().Add(duration)
	err := repo.Create(context.Background(), "Dog", "Pies", expiration)
	assert.Equalf(t, err, nil, "They should be equal")
}
