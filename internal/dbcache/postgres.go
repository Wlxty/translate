package dbcache

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type Repo struct {
	conn Connector
}

type Connector interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

func NewRepo(conn Connector) *Repo {
	return &Repo{conn}
}

func (r Repo) Read(ctx context.Context, key string) (string, time.Time, error) {
	var value string
	var ttl time.Time
	var err = r.conn.QueryRow(ctx, "SELECT value, expiration from cache WHERE key=$1", key).Scan(&value, &ttl)
	if err != nil {
		return "", ttl, err
	}
	return value, ttl, nil
}

func (r Repo) Create(ctx context.Context, key, value string, expiration time.Time) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO cache(key, value, expiration) VALUES($1, $2, $3) ON CONFLICT (key) DO UPDATE SET value=$2, expiration=$3;", key, value, expiration)
	if err != nil {
		return err
	}

	return nil
}
