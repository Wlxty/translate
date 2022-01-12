package translateapp

import (
	_ "github.com/leonelquinteros/gorand"
	"go.uber.org/zap"
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
)

//Struct of Service that got Libretranslate client and logger
type Service struct {
	Logger     *zap.SugaredLogger
	Translator Translator
}

type Translator interface {
	Languages() ([]Language, error)
	Translate(q string, source string, target string) (Word, error)
	GetLibre() libretranslate.Client
	GetCache() cache.Through
}

func NewService(logger *zap.SugaredLogger, translator Translator) *Service {
	return &Service{
		Logger:     logger,
		Translator: translator,
	}
}

// Service Translate that uses data got from LibreTranslate:5000/translate, post request. Service uses Libretranslate client.
//q = word to translate,
//source = language to translate from,
//target = language to translate to
func (service *Service) Translate(q string, source string, target string) (Word, error) {
	return service.Translator.Translate(q, source, target)
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (service *Service) Languages() ([]Language, error) {
	return service.Translator.Languages()
}
