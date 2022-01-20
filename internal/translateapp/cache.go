package translateapp

import (
	"encoding/json"
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
	value, err := c.Cache.Get(q, func() (interface{}, error) {
		return c.Libre.Translate(q, source, target)

	}, duration)
	var word Word
	json.Unmarshal([]byte(value.(string)), word)
	//Check if value.(string) is ok

	return word, err
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (c *Cache) Languages() ([]Language, error) {
	cacheKey := "languages"
	duration := time.Hour * 2
	value, err := c.Cache.Get(cacheKey, func() (interface{}, error) {
		return c.Libre.Languages()
	}, duration)
	var languages []Language
	json.Unmarshal([]byte(value.(string)), languages)
	return languages, err
}
