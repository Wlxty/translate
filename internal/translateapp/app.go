package translateapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "translateapp/internal/logger"
)

type App struct {
	Service *Service
	Router  *mux.Router
}

type Handler interface {
	ServerHTTP(syncer http.ResponseWriter, request *http.Request)
	LanguagePageHandler(writer http.ResponseWriter, request *http.Request)
	TranslatePageHandler(writer http.ResponseWriter, request *http.Request)
	HandleRequests(port string)
}

func (app App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	app.Router.ServeHTTP(writer, request)
}

func NewServer(service *Service) *App {

	router := mux.NewRouter().StrictSlash(true)

	return &App{
		Service: service,
		Router:  router,
	}
}

func (app App) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	data, err := app.Service.Languages()
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}
	app.Service.Logger.Debug("GET request on localhost:8080/languages")
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}

	fmt.Fprintf(writer, string(json))
}

func (app App) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {
	translate, _ := app.Service.Translate(request.FormValue("q"), request.FormValue("source"), request.FormValue("target"))

	app.Service.Libre.Logger.Debug("POST request on localhost:8080/translate")
	fmt.Fprintf(writer, translate)
}

func (app *App) HandleRequests(port string) {
	//create a new router
	router := app.Router

	Routes(router, app)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

func Routes(router *mux.Router, app *App) error {
	router.HandleFunc("/languages", app.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", app.TranslatePageHandler).Methods("POST")
	return nil
}
