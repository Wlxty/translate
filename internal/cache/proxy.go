package cache

import "time"

type MemoryCache interface {
	Get(key string) (bool, interface{}, error)
	Set(key string, val interface{}, expire time.Time) error
}

// NewInMemoryProxy creates an InMemory Proxy object
func NewInMemoryCache() MemoryCache {
	// TODO: implement an eviction algorithm
	return &InMemoryCache{memories: make(map[string]interface{}, 100), timeouts: make(map[string]time.Time, 100)}
}

// InMemoryProxy sets the value in the map object as a cache
type InMemoryCache struct {
	memories map[string]interface{}
	timeouts map[string]time.Time
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (p *InMemoryCache) Get(key string) (bool, interface{}, error) {
	val, ok := p.memories[key]
	expirationDate := p.timeouts[key]
	if time.Now().Sub(expirationDate) <= 0 {
		delete(p.memories, key)
		delete(p.timeouts, key)
	}
	return ok, val, nil
}

// Set sets a value to the map object as a caches
func (p *InMemoryCache) Set(key string, val interface{}, expire time.Time) error {
	_, ok := p.memories[key]
	if !ok {
		p.memories[key] = val
		p.timeouts[key] = expire
	}
	return nil
}
