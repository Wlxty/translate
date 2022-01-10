package translateapp

import (
	"github.com/leonelquinteros/gorand"
	_ "github.com/leonelquinteros/gorand"
	"go.uber.org/zap"
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
)

//Struct of Service that got Libretranslate client and logger
type Service struct {
	Logger *zap.SugaredLogger
	Libre  libretranslate.Client
	Cache  cache.Through
}

type Servicer interface {
	Languages() (string, error, string)
	Translate(q string, source string, target string) (string, error)
	GetLibre() *libretranslate.Client
	Interface() *libretranslate.Client
}

func NewService(logger *zap.SugaredLogger, libre libretranslate.Client, cache cache.Through) *Service {
	return &Service{
		Logger: logger,
		Libre:  libre,
		Cache:  cache,
	}
}

func (service *Service) Interface() *libretranslate.Client {
	var (
		wrapper Servicer = NewService(service.Logger, service.Libre, service.Cache)
	)
	libre := wrapper.GetLibre()
	return libre
}

func (service *Service) GetLibre() *libretranslate.Client {
	return &service.Libre
}

// Service Translate that uses data got from LibreTranslate:5000/translate, post request. Service uses Libretranslate client.
//q = word to translate,
//source = language to translate from,
//target = language to translate to
func (service *Service) Translate(q string, source string, target string) (string, error) {
	var (
		libre libretranslate.Libre = service.Interface()
	)
	return libre.Translate(q, source, target)
}

func (service *Service) GetLanguages() (interface{}, error) {
	var (
		libre libretranslate.Libre = service.Interface()
	)
	return libre.Languages()
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (service *Service) Languages() (string, error, string) {
	cacheKey, _ := gorand.GetAlphaNumString(24)
	value, err := service.Cache.Get(cacheKey, service.GetLanguages)
	return value.(string), err, cacheKey
}
