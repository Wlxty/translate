package translateapp

import (
	"encoding/json"
	"go.uber.org/zap"
	"translateapp/internal/libretranslate"
)

type Service struct {
	Logger *zap.SugaredLogger
	Libre  libretranslate.Client
}

type Servicer interface {
	Languages() (string, error)
	Translate(q string, source string, target string) (string, error)
}

func (service *Service) Languages() ([]Language, error) {
	data, err := service.Libre.Languages()
	languages := []Language{}
	json.Unmarshal([]byte(data), &languages)
	return languages, err
}

func (service *Service) Translate(q string, source string, target string) (string, error) {
	return service.Libre.Translate(q, source, target)
}
