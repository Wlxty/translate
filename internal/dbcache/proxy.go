package dbcache

import (
	"context"
	"fmt"
	"log"
	"time"
)

// InMemoryProxy sets the value in the map object as a cache
type DBCache struct {
	repo Repository
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (dbc *DBCache) Get(key string) (bool, interface{}, error) {
	value, err := dbc.repo.Read(context.Background(), key)
	if err == nil {
		return true, value, nil
	}
	return false, nil, err
}

// Set sets a value to the map object as a caches
func (dbc *DBCache) Set(key string, val interface{}, expire time.Duration) error {
	// Executing SQL query for insertion
	log.Println("Inserting to database")

	expiration := time.Now().Add(expire)
	data, ok := val.(string)
	if !ok {
		return fmt.Errorf("invalid conversion in set method")
	}
	return dbc.repo.Create(context.Background(), key, data, expiration)
}
