package server

import(
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"translateapp/internal/logging"
)
func (ts *Server) LanguagePageHandler(w http.ResponseWriter, req *http.Request) {
	data := ts.languages
	jsonify, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w,"Error: %s", err.Error())
		return
	}
	logging.NewLogger("https://localhost:8080/languages", "GET")

	fmt.Fprintf(w,string(jsonify))
}

func (ts *Server) TranslatePageHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"word": req.FormValue("word"),
		"source": req.FormValue("source"),
		"target": req.FormValue("target"),
	}

	if data["word"] == "" || data["source"] == "" || data["target"] == "" {
		fmt.Fprintf(w, "Error: %s", errors.New("Require all parameters to be filled"))
		return
	}
	_, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w,"Error: %s", err.Error())
		return
	}

	repository := ts.translated
	output := repository.Translate()
	JsonOutput, _ := json.Marshal(output)
	logging.NewLogger("https://localhost:8080/languages", "POST")
	fmt.Fprintf(w, string(JsonOutput))
}

