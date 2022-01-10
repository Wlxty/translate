package translateapp

import (
	_ "github.com/leonelquinteros/gorand"
	"go.uber.org/zap"
)

//Struct of Service that got Libretranslate client and logger
type Service struct {
	Logger *zap.SugaredLogger
	Cached Cacher
}

type Servicer interface {
	Languages() (string, error, string)
	Translate(q string, source string, target string) (string, error)
}

func NewService(logger *zap.SugaredLogger, cacher Cacher) *Service {
	return &Service{
		Logger: logger,
		Cached: cacher,
	}
}

// Service Translate that uses data got from LibreTranslate:5000/translate, post request. Service uses Libretranslate client.
//q = word to translate,
//source = language to translate from,
//target = language to translate to
func (service *Service) Translate(q string, source string, target string) (string, error) {
	return service.Cached.Translate(q, source, target)
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (service *Service) Languages() (string, error, string) {
	return service.Cached.Languages()
}
