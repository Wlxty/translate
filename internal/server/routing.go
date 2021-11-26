package server

import(
	"net/http"
	"fmt"
)
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

func TranslateHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/translate" {
		// Request is plain "/task/", without trailing ID.
		if req.Method == http.MethodPost {
			data := map[string]string{
				"text": "Hello World",
			}
			renderJSON(w, data)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}
