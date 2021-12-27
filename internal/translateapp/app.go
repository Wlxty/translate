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

// Starting Http server on gorilla mux router
func (app *App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	app.Router.ServeHTTP(writer, request)
}

// Starting application
func NewApp(service *Service) *App {

	router := mux.NewRouter().StrictSlash(true)

	return &App{
		Service: service,
		Router:  router,
	}
}

// Request to fetch all languages from Libretranslate service.
func (app *App) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	data, err := app.Service.Languages()
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}
	json, err := json.Marshal(data)
	app.Service.Logger.Debug("GET request on localhost:8080/languages")
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}
	fmt.Fprintf(writer, "%s", string(json))
}

// Request to get translation from Libretranslate service.
func (app *App) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {
	translate, _ := app.Service.Translate(request.FormValue("q"), request.FormValue("source"), request.FormValue("target"))

	app.Service.Libre.Logger.Debug("POST request on localhost:8080/translate")
	fmt.Fprintf(writer, "%s", translate)
}

// Method to handle all requests
func (app *App) HandleRequests(port string) {
	//create a new router
	router := app.Router

	Routes(router, app)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

// Routing,
//if you want to add new route, add it here
func Routes(router *mux.Router, app *App) {
	router.HandleFunc("/languages", app.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", app.TranslatePageHandler).Methods("POST")
}
