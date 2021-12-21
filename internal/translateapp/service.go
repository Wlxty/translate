package translateapp

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Service struct{}

func (service *Service) Languages() ([]Language, error) {
	var languages Language
	return languages.Languages(), nil
}

func (service *Service) Translate(data url.Values, url string) (string, error) {
	resp, err := http.PostForm(url+"translate", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	jsonify, err := json.Marshal(res)

	return string(jsonify), err
}
