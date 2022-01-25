package dbcache

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

type Repo struct {
	conn *pgx.Conn
}

func NewRepo(conn *pgx.Conn) *Repo {
	return &Repo{conn}
}

func (r Repo) Read(ctx context.Context, key string) (string, error) {
	var value string
	err := r.conn.QueryRow(ctx, "SELECT value from cache WHERE key=$1", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (r Repo) Create(ctx context.Context, key, value string, expiration time.Time) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO cache(key, value, expiration) VALUES($1, $2, $3)", key, value, expiration)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
