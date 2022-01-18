package translateapp

import (
	"encoding/json"
	"translateapp/internal/libretranslate"
)

type LibreWrapper struct {
	libre *libretranslate.Client
}

func NewLibreWrapper(libre *libretranslate.Client) *LibreWrapper {
	return &LibreWrapper{libre: libre}
}

func (lw *LibreWrapper) Languages() ([]Language, error) {
	languages, err := lw.libre.Languages()
	data := []Language{}
	json.Unmarshal([]byte(languages), &data)
	return data, err
}
func (lw *LibreWrapper) Translate(q string, source string, target string) (Word, error) {
	translate, err := lw.libre.Translate(q, source, target)
	data := Word{}
	json.Unmarshal([]byte(translate), &data)
	return data, err
}
