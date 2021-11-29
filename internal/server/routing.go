package server

import(
	"net/http"
	"fmt"
	"encoding/json"
)
func (ts *taskServer) LanguagePageHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/languages/" {
		// Request is plain "/task/", without trailing ID.
		if req.Method == http.MethodGet {
			ts.getArrayLanguages(w, req)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (ts *taskServer) HomePageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, ":Endpoint called: HomePage")
}

func (ts *taskServer) TranslatePageHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"word": req.FormValue("word"),
		"source": req.FormValue("source"),
		"target": req.FormValue("target"),
	}
	jsonify, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w,"Error: %s", err.Error())
	} else {
		fmt.Fprintf(w,string(jsonify))
	}
}

