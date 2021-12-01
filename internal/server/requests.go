package server

import(
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"translateapp/internal/languages"
)
func (ts *taskServer) LanguagePageHandler(w http.ResponseWriter, req *http.Request) {
	repository := languages.New()
	data := repository.Languages()
	jsonify, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w,"Error: %s", err.Error())
		return
	}
	fmt.Fprintf(w,string(jsonify))
}

func (ts *taskServer) TranslatePageHandler(w http.ResponseWriter, req *http.Request) {
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

	output := map[string]string{
		"TranslatedWord": "Translated word",
	}
	JsonOutput, _ := json.Marshal(output)

	fmt.Fprintf(w, string(JsonOutput))
}

