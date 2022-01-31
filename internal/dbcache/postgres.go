package dbcache

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Repo struct {
	conn *pgx.Conn
}

func NewRepo(conn *pgx.Conn) *Repo {
	return &Repo{conn}
}

func (r Repo) Read(ctx context.Context, key string) (string, time.Time, error) {
	var value string
	var ttl time.Time
	err := r.conn.QueryRow(ctx, "SELECT value, expiration from cache WHERE key=$1", key).Scan(&value, &ttl)
	if err != nil {
		return "", ttl, err
	}
	return value, ttl, nil
}

func (r Repo) Update(ctx context.Context, key, value string, expiration time.Time) error {
	err := r.conn.QueryRow(ctx, "UPDATE cache SET value=$1, expiration=$2 WHERE key=$3", value, expiration, key)
	if err != nil {
		return fmt.Errorf("Cannot update such row.")
	}
	return nil
}

func (r Repo) Create(ctx context.Context, key, value string, expiration time.Time) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO cache(key, value, expiration) VALUES($1, $2, $3)", key, value, expiration)
	if err != nil {
		return err
	}

	return nil
}
