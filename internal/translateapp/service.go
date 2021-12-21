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

func (service *Service) Languages() ([]Language, error) {
	var languages Language
	return languages.Languages(), nil
}

func (service *Service) Translate(data url.Values) (string, error) {
	resp, err := http.PostForm(service.Libre.Host+"translate", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	jsonify, err := json.Marshal(res)

	return string(jsonify), err
}
