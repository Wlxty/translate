package translateapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/url"
	"translateapp/internal/libretranslate"
	_ "translateapp/internal/logger"
)

type Server struct {
	Service *Service
	Router  *mux.Router
	Logger  *zap.SugaredLogger
	Libre   libretranslate.Client
}

func NewServer(client libretranslate.Client, logger *zap.SugaredLogger) *Server {

	router := mux.NewRouter().StrictSlash(true)
	var service Service

	return &Server{
		Service: &service,
		Router:  router,
		Logger:  logger,
		Libre:   client,
	}
}

func (server Server) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	data, err := server.Service.Languages()

	server.Logger.Debug("GET request on localhost:8080/languages")
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}

	fmt.Fprintf(writer, string(json))
}

func (server Server) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {
	data := url.Values{
		"q":      {request.FormValue("q")},
		"source": {request.FormValue("source")},
		"target": {request.FormValue("target")},
	}
	translate, _ := server.Service.Translate(data, server.Libre.Host)

	server.Libre.Logger.Debug("POST request on localhost:8080/translate")
	fmt.Fprintf(writer, translate)
}

func (server *Server) HandleRequests(port string) {
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
