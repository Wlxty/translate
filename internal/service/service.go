package service

import (
	"encoding/json"
	"fmt"
	tr "github.com/snakesel/libretranslate"
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
	translate := tr.New(tr.Config{
		Url: "172.19.0.3:5000",
		Key: "XXX",
	})
	output := word.Translate()
	txt, _ := translate.Translate(output.TranslatedWord, "pl", "en")
	fmt.Println(txt)
	output = models.Word{txt}
	jsonify, err := json.Marshal(output)
	if err != nil {
		fmt.Fprintf(writer, "Error: %w", err)
		return nil
	}

	fmt.Fprintf(writer, string(jsonify))
	return nil
}
