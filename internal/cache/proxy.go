package cache

import "time"

// NewInMemoryProxy creates an InMemory Proxy object
func NewInMemoryCache() *InMemoryCache {
	// TODO: implement an eviction algorithm
	return &InMemoryCache{
		data:           make(map[string]interface{}),
		expirationDate: make(map[string]time.Time),
	}
}

// InMemoryProxy sets the value in the map object as a cache
type InMemoryCache struct {
	data           map[string]interface{}
	expirationDate map[string]time.Time
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (p *InMemoryCache) Get(key string) (bool, interface{}, error) {
	val, ok := p.data[key]
	expirationDate := p.expirationDate[key]
	if time.Now().Sub(expirationDate) <= 0 {
		delete(p.data, key)
		delete(p.expirationDate, key)
	}
	return ok, val, nil
}

// Set sets a value to the map object as a caches
func (p *InMemoryCache) Set(key string, val interface{}, expire time.Duration) error {
	p.data[key] = val
	p.expirationDate[key] = time.Now().Add(expire)
	return nil
}
