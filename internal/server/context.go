package server
import(
	"net/http"
	"log"
	"translateapp/internal/languages"
)

type taskServer struct {
	store *languages.Repository
}

func NewTaskServer() *taskServer {
	repository := languages.New()
	return &taskServer{store: &repository}
}

func (ts *taskServer) getArrayLanguages(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all languages at %s\n", req.URL.Path)

	languages := ts.store.Languages()
	renderJSON(w, languages)
}
