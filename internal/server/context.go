package server
import(
	"translateapp/internal/models"
        "github.com/gorilla/mux"
)

type Server struct {
	languages *models.LanguageRepository
	translated *models.TranslatedWordRepository
	router *mux.Router
}

func NewServer() *Server {
	lang := models.NewLanguage()
	word := models.NewTranslation("word")
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		languages: &lang,
		translated: &word,
		router: router,
	}
}
