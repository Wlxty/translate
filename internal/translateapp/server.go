package translateapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "translateapp/internal/logger"
)

type Server struct {
	Service *Service
	Router  *mux.Router
}

func NewServer(service *Service) *Server {

	router := mux.NewRouter().StrictSlash(true)

	return &Server{
		Service: service,
		Router:  router,
	}
}

func (server Server) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	data, err := server.Service.Languages()
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}
	server.Service.Logger.Debug("GET request on localhost:8080/languages")
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}

	fmt.Fprintf(writer, string(json))
}

func (server Server) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {
	translate, _ := server.Service.Translate(request.FormValue("q"), request.FormValue("source"), request.FormValue("target"))

	server.Service.Libre.Logger.Debug("POST request on localhost:8080/translate")
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
