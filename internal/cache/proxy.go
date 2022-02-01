package cache

import (
	"go.uber.org/zap"
	"time"
)

// NewInMemoryProxy creates an InMemory Proxy object
func NewInMemoryCache(logger *zap.SugaredLogger) *InMemoryCache {
	// TODO: implement an eviction algorithm
	return &InMemoryCache{
		data:           make(map[string]interface{}),
		expirationDate: make(map[string]time.Time),
		logger:         logger,
	}
}

// InMemoryProxy sets the value in the map object as a cache
type InMemoryCache struct {
	data           map[string]interface{}
	expirationDate map[string]time.Time
	logger         *zap.SugaredLogger
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (mc *InMemoryCache) Get(key string) (bool, interface{}, error) {
	val, ok := mc.data[key]
	ttl := mc.expirationDate[key]
	expiration := ttl.Sub(time.Now())
	mc.logger.Debugf("Expired in %s", expiration)
	if expiration <= 0 {
		delete(mc.data, key)
		delete(mc.expirationDate, key)
		mc.logger.Debugf("Row has been deleted")

	}
	return ok, val, nil
}

// Set sets a value to the map object as a caches
func (mc *InMemoryCache) Set(key string, val interface{}, expire time.Duration) error {
	mc.data[key] = val
	mc.expirationDate[key] = time.Now().Add(expire)
	return nil
}
