package dbcache

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

// InMemoryProxy sets the value in the map object as a cache
type DBCache struct {
	repo   Repository
	logger *zap.SugaredLogger
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (dbc *DBCache) Get(key string) (bool, interface{}, error) {
	value, ttl, err := dbc.repo.Read(context.Background(), key)
	if err == nil {
		expiration := ttl.Sub(time.Now())
		if expiration <= 0 {
			return false, nil, err
		}
		dbc.logger.Infof("Getting row from database")
		return true, value, nil
	}
	return false, nil, err
}

// Set sets a value to the map object as a caches
func (dbc *DBCache) Set(key string, val interface{}, expire time.Duration) error {
	// Executing SQL query for insertion
	dbc.logger.Infof("Inserting to database")
	expiration := time.Now().Add(expire)
	data, ok := val.(string)
	if !ok {
		return fmt.Errorf("invalid conversion in set method")
	}
	return dbc.repo.Create(context.Background(), key, data, expiration)
}
