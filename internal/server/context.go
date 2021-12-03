package server
import(
	"translateapp/internal/models"
)

type Server struct {
	languages *models.LanguageRepository
	translated *models.TranslatedWordRepository
}

func NewServer() *Server {
	lang := models.NewLanguage()
	word := models.NewTranslation("word")
	return &Server{
		languages: &lang,
		translated: &word,
	}
}
