package translateapp

import (
	"encoding/json"
	"time"
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
)

type Cache struct {
	Libre libretranslate.Libre
	Cache cache.Through
}

func NewCache(libre libretranslate.Libre, cache cache.Through) Cache {
	return Cache{libre, cache}
}
func (c *Cache) GetCache() cache.Through {
	return c.Cache
}

type Cacher interface {
	Translate(q string, source string, target string) (Word, error)
	GetLanguages() (interface{}, error)
	Languages() ([]Language, error)
	GetLibre() libretranslate.Client
	GetCache() cache.Through
}

func (c *Cache) GetLibre() libretranslate.Client {
	return *c.Libre.GetLibre()
}

func (c *Cache) GetCatche() cache.Through {
	return c.GetCatche()
}

func (c *Cache) Translate(q string, source string, target string) (Word, error) {
	duration := time.Hour * 2
	value, err := c.Cache.Get(q, func() (interface{}, error) {
		var translator, _ = c.Libre.Translate(q, source, target)
		return translator, nil

	}, duration)
	var word Word
	json.Unmarshal([]byte(value.(string)), word)
	//Check if value.(string) is ok

	return word, err
}

func (c *Cache) GetLanguages() (interface{}, error) {
	return c.Libre.Languages()
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (c *Cache) Languages() ([]Language, error) {
	cacheKey := "languages"
	duration := time.Hour * 2
	value, err := c.Cache.Get(cacheKey, c.GetLanguages, duration)
	var languages []Language
	json.Unmarshal([]byte(value.(string)), languages)
	return languages, err
}
