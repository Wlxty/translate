package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"translateapp/internal/translateapp/models"
)

type Service struct{}

func (service *Service) Languages(writer http.ResponseWriter) error {
	var languages models.Language
	output := languages.Languages()
	jsonify, err := json.Marshal(output)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
		return nil
	}

	fmt.Fprintf(writer, string(jsonify))
	return nil
}

func (service *Service) Translate(writer http.ResponseWriter) error {
	var word models.Word
	output := word.Translate()
	jsonify, err := json.Marshal(output)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
		return nil
	}

	fmt.Fprintf(writer, string(jsonify))
	return nil
}
