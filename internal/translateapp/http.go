package translateapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	s "translateapp/internal/server"
)

type Service struct {
	server *s.Server
}

func (service Service) TranslatePageHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"word":   req.FormValue("word"),
		"source": req.FormValue("source"),
		"target": req.FormValue("target"),
	}

	if data["word"] == "" || data["source"] == "" || data["target"] == "" {
		fmt.Fprintf(w, "Error: %s", errors.New("Require all parameters to be filled"))
		return
	}
	_, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	translated := service.server.Translated
	output := translated.Translate()
	JsonOutput, _ := json.Marshal(output)

	fmt.Fprintf(w, string(JsonOutput))

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("POST request on localhost:8080/translate")
}

func (service Service) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	var data = service.server.Languages
	jsonify, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(writer, string(jsonify))

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("GET request on localhost:8080/languages")
}

func HandleRequests(port string) {
	server := s.NewServer()

	//create a new router
	router := server.Router

	request := Service{server}
	Routes(router, &request)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

func Routes(router *mux.Router, service *Service) error {
	router.HandleFunc("/languages", service.LanguagePageHandler).Methods("GET")
	router.HandleFunc(
		"/translate",
		service.TranslatePageHandler,
	).Methods("POST")
	return nil
}
