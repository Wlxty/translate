package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"translateapp/internal/translateapp"
)

type Service struct{}

func (service *Service) Languages(writer http.ResponseWriter) error {
	var languages translateapp.Language
	output := languages.Languages()
	jsonify, err := json.Marshal(output)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
		return nil
	}

	fmt.Fprintf(writer, string(jsonify))
	return nil
}

func (service *Service) Translate(writer http.ResponseWriter, data url.Values) error {
	resp, err := http.PostForm("http://libretranslate:5000//translate", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	jsonify, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
		fmt.Fprintf(writer, "Error: %w", err)
		return nil
	}

	fmt.Fprintf(writer, string(jsonify))
	return nil
}
