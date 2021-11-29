package server
import(
	"net/http"
	"log"
)
func (ts *taskServer) getArrayLanguages(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all languages at %s\n", req.URL.Path)

	languages := ts.store.Languages()
	renderJSON(w, languages)
}
