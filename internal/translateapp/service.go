package translateapp

import (
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/url"
	"translateapp/internal/libretranslate"
)

type Service struct {
	Logger *zap.SugaredLogger
	Libre  libretranslate.Client
}

type Servicer interface {
	Languages() ([]Language, error)
	Translate(q string, source string, target string) (string, error)
}

func (service *Service) Languages() ([]Language, error) {
	var languages Language
	return languages.Languages(), nil
}

func (service *Service) Translate(q string, source string, target string) (string, error) {
	data := url.Values{
		"q":      {q},
		"source": {source},
		"target": {target},
	}
	resp, err := http.PostForm(service.Libre.Host+"translate", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	jsonify, err := json.Marshal(res)

	return string(jsonify), err
}
