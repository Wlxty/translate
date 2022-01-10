package cache

import "time"

type Proxy interface {
	Get(key string) (bool, interface{}, error)
	Set(key string, val interface{}) error
}

// NewInMemoryProxy creates an InMemory Proxy object
func NewInMemoryProxy() Proxy {
	// TODO: implement an eviction algorithm
	return &InMemoryProxy{m: make(map[string]interface{}, 100)}
}

// InMemoryProxy sets the value in the map object as a cache
type InMemoryProxy struct {
	m map[string]struct{
		value interface{}
		ttl time.Time
	}
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (p *InMemoryProxy) Get(key string) (bool, interface{}, error) {
	val, ok := p.m[key]
	if time.Now() > val.ttl {
		delete(p.m[key])
		return false, nul, nil
	}
	return ok, val.value, nil
}

// Set sets a value to the map object as a caches
func (p *InMemoryProxy) Set(key string, val interface{}, timeout *time.Duration) error {
	p.m[key] = struct{
		val: val,
		ttl: time.Now() + *titimeout,
	}
	return nil
}
