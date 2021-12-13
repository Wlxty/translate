package server

import (
	"github.com/gorilla/mux"
	models "translateapp/internal/translateapp/models"
)

type Server struct {
	Languages  []models.Language
	Translated models.Word
	Router     *mux.Router
}

func NewServer() *Server {
	word := models.NewWord("word")
	var language models.Language
	lang := language.Languages()
	translate := word.Translate()

	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		Languages:  lang,
		Translated: translate,
		Router:     router,
	}
}
