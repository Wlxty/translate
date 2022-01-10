package translateapp

import (
	"fmt"
	"strings"
	"time"
)

type Cache struct {
	cache      Cacher
	translator Translator
}

// Cacher defines interface for cache implementers.
type Cacher interface {
	// Set adds value to the cache under specified key.
	// The timeout argument is optional and defaults to predefined timeout.
	// It's the number of seconds the value should be stored in the cache.
	// Passing in nil for timeout will cache the value forever.
	// A timeout of 0 won't cache the value.
	Set(key string, value interface{}, timeout *time.Duration) error
	// Get retrieves data from the cache.
	// If the data doesn't exist in the cache, Get() returns ErrDoesNotExist.
	Get(key string) (interface{}, error)
}

// translator := libretranslate.Client
// cachedTranslator := NewCache(NewInMemoryProxy(), translator)
// cachedTranslator := NewCache(NewInDBProxy(), translator)
// NewService(logger, cachedTranslator) // with cache
// NewService(logger, translator) // original
func NewCache(cache Cacher, translator Translator) Cache {
	return Cache{
		cache:      cache,
		translator: translator,
	}
}

// Translate returns translated text.
func (c *Cache) Translate(text, from, to string) (string, error) {
	val, err := c.cache.Get(strings.ToLower(text))
	if err != nil {
		return "", err
	}
	translatedText, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("casting error")
	}
	return translatedText, nil
}

func (c *Cache) Languages() ([]Language, error) {
	val, err := c.cache.Get(text)
	if err != nil {
		return nil, err
	}
	languages, ok := val.([]Language)
	if !ok {
		return nil, fmt.Errorf("casting error")
	}
	return languages, nil
}
