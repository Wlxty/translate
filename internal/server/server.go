package server

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/url"
	"translateapp/internal/service"
)

type Server struct {
	Service *service.Service
	Router  *mux.Router
	Logger  *zap.Logger
}

func NewServer() *Server {

	router := mux.NewRouter().StrictSlash(true)
	var service service.Service
	logger, _ := zap.NewProduction()
	return &Server{
		Service: &service,
		Router:  router,
		Logger:  logger,
	}
}

func (server Server) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	server.Service.Languages(writer)
	logger := server.Logger
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("GET request on localhost:8080/languages")
}

func (server Server) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {

	data := url.Values{
		"q":      {request.FormValue("q")},
		"source": {request.FormValue("source")},
		"target": {request.FormValue("target")},
	}
	server.Service.Translate(writer, data)

	logger := server.Logger
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("POST request on localhost:8080/translate")
}

func HandleRequests(port string) {
	server := NewServer()

	//create a new router
	router := server.Router

	Routes(router, server)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

func Routes(router *mux.Router, server *Server) error {
	router.HandleFunc("/languages", server.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", server.TranslatePageHandler).Methods("POST")
	return nil
}
