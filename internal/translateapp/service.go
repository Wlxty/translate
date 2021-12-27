package translateapp

import (
	"encoding/json"
	"go.uber.org/zap"
	"translateapp/internal/libretranslate"
)

//Struct of Service that got Libretranslate client and logger
type Service struct {
	Logger *zap.SugaredLogger
	Libre  libretranslate.Client
}

type Servicer interface {
	Languages() ([]Language, error)
	Translate(q string, source string, target string) (Word, error)
}

// Service languages that uses data got from LibreTranslate:5000/languages, get request. Service uses Libretranslate client.
func (service *Service) Languages() ([]Language, error) {
	data, err := service.Libre.Languages()
	languages := []Language{}
	json := json.Unmarshal([]byte(data), &languages)
	if json != nil {
		service.Logger.Debug("Service Languages: Not valid Json")
	}
	service.Logger.Debug("Service Languages works fine")
	return languages, err
}

// Service Translate that uses data got from LibreTranslate:5000/translate, post request. Service uses Libretranslate client.
//q = word to translate,
//source = language to translate from,
//target = language to translate to
func (service *Service) Translate(q string, source string, target string) (string, error) {
	return service.Libre.Translate(q, source, target)
}
