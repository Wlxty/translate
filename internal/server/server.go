package server
import(
	"translateapp/internal/languagesstore"
	"net/http"
	"fmt"
)
type taskServer struct {
	store *languagesstore.LanguagesStore
}

func NewTaskServer() *taskServer {
	store := languagesstore.New()
	return &taskServer{store: &store}
}

func (ts *taskServer) TaskHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/languages/" {
		// Request is plain "/task/", without trailing ID.
		if req.Method == http.MethodGet {
			ts.getAllLanguagesHandler(w, req)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}


