package translateapp

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Cache struct {
	Libre Translator
	Cache Cacher
}

type Cacher interface {
	Get(key string, req func() (interface{}, error), expire time.Duration) (interface{}, error)
}

func NewCache(libre Translator, cache Cacher) *Cache {
	return &Cache{libre, cache}
}

func Conv(input interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (c *Cache) Translate(q string, source string, target string) (Word, error) {
	duration := time.Hour * 2
	data, err := c.Cache.Get(q, func() (interface{}, error) {
		translation, err := c.Libre.Translate(q, source, target)
		return Conv(translation, err)

	}, duration)
	str, ok := data.(string)

	if !ok {
		return Word{}, fmt.Errorf("invalid conversion")
	}
	log.Println(str)
	var value Word
	err = json.Unmarshal([]byte(str), &value)
	if err != nil {
		return Word{}, fmt.Errorf("invalid json")
	}
	return value, nil
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (c *Cache) Languages() ([]Language, error) {
	cacheKey := "languages"
	duration := time.Hour * 2
	data, err := c.Cache.Get(cacheKey, func() (interface{}, error) {
		languages, err := c.Libre.Languages()
		return Conv(languages, err)

	}, duration)
	str, ok := data.(string)

	if !ok {
		return nil, fmt.Errorf("invalid conversion")
	}
	log.Println(str)
	var value []Language
	err = json.Unmarshal([]byte(str), &value)
	if err != nil {
		return nil, fmt.Errorf("invalid json")
	}
	return value, err
}
