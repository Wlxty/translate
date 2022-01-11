package cache

import "time"

type Proxy interface {
	Get(key string) (bool, interface{}, error)
	Set(key string, val interface{}, expire time.Time) error
}

// NewInMemoryProxy creates an InMemory Proxy object
func NewInMemoryProxy() Proxy {
	// TODO: implement an eviction algorithm
	return &InMemoryProxy{memories: make(map[string]interface{}, 100), timeouts: make(map[string]time.Time, 100)}
}

// InMemoryProxy sets the value in the map object as a cache
type InMemoryProxy struct {
	memories map[string]interface{}
	timeouts map[string]time.Time
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (p *InMemoryProxy) Get(key string) (bool, interface{}, error) {
	val, ok := p.memories[key]
	expirationDate := p.timeouts[key]
	if time.Now().Sub(expirationDate) <= 0 {
		delete(p.memories, key)
		delete(p.timeouts, key)
	}
	return ok, val, nil
}

// Set sets a value to the map object as a caches
func (p *InMemoryProxy) Set(key string, val interface{}, expire time.Time) error {
	p.memories[key] = val
	p.timeouts[key] = expire
	return nil
}
