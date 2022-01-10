package translateapp

import (
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
)

type Cache struct {
	Libre libretranslate.Client
	Cache cache.Through
}

type Cacher interface {
	Translate(q string, source string, target string) (string, error)
	GetLanguages() (interface{}, error)
	Languages() (string, error, string)
	GetLibre() libretranslate.Client
	GetCache() cache.Through
}

func (c *Cache) GetLibre() libretranslate.Client {
	return c.Libre
}

func (c *Cache) GetCatche() cache.Through {
	return c.Cache
}

func (c *Cache) Translate(q string, source string, target string) (string, error) {
	return c.Libre.Translate(q, source, target)
}

func (c *Cache) GetLanguages() (interface{}, error) {
	return c.Libre.Languages()
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (c *Cache) Languages() (string, error, string) {
	cacheKey := "languages"
	value, err := c.Cache.Get(cacheKey, c.GetLanguages)
	return value.(string), err, cacheKey
}
