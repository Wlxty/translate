package dbcache

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

// NewInMemoryProxy creates an InMemory Proxy object
func NewInDBCache(connection pgx.Conn) *DBCache {
	// TODO: implement an eviction algorithm
	return &DBCache{
		connection,
	}
}

// InMemoryProxy sets the value in the map object as a cache
type DBCache struct {
	Connection pgx.Conn
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (dbc *DBCache) Get(key string) (bool, interface{}, error) {
	var value string
	var expiration string
	err := dbc.Connection.QueryRow(context.Background(), "SELECT value, expiration from cache WHERE key=$1", key).Scan(&value, &expiration)
	if err == nil {
		fmt.Println("Error occur while finding cache: ", err)
		return false, nil, nil
	}
	return true, value, err
}

// Set sets a value to the map object as a caches
func (dbc *DBCache) Set(key string, val interface{}, expire time.Duration) error {
	// Executing SQL query for insertion
	expiration := time.Now().Add(expire).Unix()
	_, err := dbc.Connection.Exec(context.Background(), "INSERT INTO cache(key, value, expiration) VALUES($1, $2, $3)", key, val, expiration)
	if err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return nil
	}
	return err
}
