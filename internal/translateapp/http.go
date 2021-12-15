package translateapp

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/url"
	s "translateapp/internal/server"
)

type Client struct {
	server *s.Server
}

func (client Client) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	client.server.Service.Languages(writer)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("GET request on localhost:8080/languages")
}

func (client Client) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {

	data := url.Values{
		"q":      {request.FormValue("q")},
		"source": {request.FormValue("source")},
		"target": {request.FormValue("target")},
	}
	client.server.Service.Translate(writer, data)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("POST request on localhost:8080/translate")
}

func HandleRequests(port string) {
	server := s.NewServer()

	//create a new router
	router := server.Router

	request := Client{server}
	Routes(router, &request)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

func Routes(router *mux.Router, client *Client) error {
	router.HandleFunc("/languages", client.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", client.TranslatePageHandler).Methods("POST")
	return nil
}
