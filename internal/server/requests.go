package server

import(
	"net/http"
	"fmt"
	"errors"
	"log"
	"encoding/json"
)
func (ts *Server) LanguagePageHandler(w http.ResponseWriter, req *http.Request) {
	repository := ts.languages
	data := repository.Languages()
	jsonify, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w,"Error: %s", err.Error())
		return
	}
	log.Println("request type: GET, endpoint: localhost:8080/languages")
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
	output := repository.TranslatedWord()
	JsonOutput, _ := json.Marshal(output)
	log.Println("request type: POST, endpoint: localhost:8080/translate, variables: {", data["word"], data["source"], data["target"], "}")

	fmt.Fprintf(w, string(JsonOutput))
}

