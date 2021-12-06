package server
import(
	models "translateapp/internal/translate"
        "github.com/gorilla/mux"
)

type Server struct {
	languages []models.Language
	translated models.Word
	router *mux.Router
}

func NewServer() *Server {
	word := models.NewWord("word")
	var language models.Language
	lang := language.Languages()
	translate := word.Translate()

	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		languages: lang,
		translated: translate,
		router: router,
	}
}
