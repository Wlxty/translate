package translateapp

import (
	"encoding/json"
	"translateapp/internal/libretranslate"

	"go.uber.org/zap"
)

type Translator interface {
	// Translate returns translated text.
	Translate(text, from, to string) (string, error)
	Languages() ([]Language, error)
}

//Struct of Service that got Libretranslate client and logger
type Service struct {
	Logger *zap.SugaredLogger
	// Libre  libretranslate.Client
	Libre Translator
}

func NewService(logger *zap.SugaredLogger, libre Translator) *Service {
	return &Service{
		Logger: logger,
		Libre:  libre,
	}
}

func (cache *Service) GetLibre() *libretranslate.Client {
	return &cache.Libre
}

func (cache *Service) Interface() *libretranslate.Client {
	var (
		wrapper Servicer = NewService(cache.Logger, cache.Libre)
	)
	libre := wrapper.GetLibre()
	return libre
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (cache *Service) Languages() ([]Language, error) {
	var (
		libre libretranslate.Libre = cache.Interface()
	)
	data, err := libre.Languages()

	var languages []Language
	json := json.Unmarshal([]byte(data), &languages)
	if json != nil {
		cache.Logger.Debug("Service Languages: Not valid Json")
	}
	cache.Logger.Debug("Service Languages works fine")
	return languages, err
}

// Service Translate that uses data got from LibreTranslate:5000/translate, post request. Service uses Libretranslate client.
//q = word to translate,
//source = language to translate from,
//target = language to translate to
func (cache *Service) Translate(q string, source string, target string) (string, error) {
	var (
		libre libretranslate.Libre = cache.Interface()
	)
	return libre.Translate(q, source, target)
}
