package translateapp

import (
	"time"
	"translateapp/internal/cache"
)

type Cache struct {
	Libre Translator
	Cache cache.Through
}

func NewCache(libre Translator, cache cache.Through) *Cache {
	return &Cache{libre, cache}
}

func (c *Cache) GetCache() cache.Through {
	return c.Cache
}

type Cacher interface {
	Translate(q string, source string, target string) (Word, error)
	Languages() ([]Language, error)
}

func (c *Cache) Translate(q string, source string, target string) (Word, error) {
	duration := time.Hour * 2
	translation, err := c.Libre.Translate(q, source, target)
	_, err = c.Cache.Get(q, func() (interface{}, error) {
		return translation, err

	}, duration)
	return translation, err
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (c *Cache) Languages() ([]Language, error) {
	cacheKey := "languages"
	duration := time.Hour * 2
	value, err := c.Cache.Get(cacheKey, func() (interface{}, error) {
		return c.Libre.Languages()
	}, duration)
	return value.([]Language), err
}
