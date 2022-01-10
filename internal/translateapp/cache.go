package translateapp

import (
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
)

type Cache struct {
	Libre libretranslate.Libre
	Cache cache.Through
}

func (c *Cache) GetCache() cache.Through {
	return c.Cache
}

type Cacher interface {
	Translate(q string, source string, target string) (string, error)
	GetLanguages() (interface{}, error)
	Languages() (string, error, string)
	GetLibre() libretranslate.Client
	GetCache() cache.Through
}

func (c *Cache) GetLibre() libretranslate.Client {
	return *c.Libre.GetLibre()
}

func (c *Cache) GetCatche() cache.Through {
	return c.GetCatche()
}

func (c *Cache) Translate(q string, source string, target string) (string, error) {
	value, err := c.Cache.Get(q, func() (interface{}, error) {
		var translator, _ = c.Libre.Translate(q, source, target)
		return translator, nil

	})
	return value.(string), err
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
