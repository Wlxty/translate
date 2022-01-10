package translateapp

import (
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
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
	LanguagePageHandler(writer http.ResponseWriter, request *http.Request)
	TranslatePageHandler(writer http.ResponseWriter, request *http.Request)
	HandleRequests(port string)
	Routes(router *mux.Router)
	GetRouter() *mux.Router
	GetService() *Service
}

func (app *App) GetRouter() *mux.Router {
	return app.Router
}

func (app *App) GetService() *Service {
	return app.Service
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
	data, err, cache := app.GetService().Languages()
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}
	app.Service.Logger.Debug("GET request on localhost:8080/languages")
	_, cached, _ := app.Service.Cached.Cache.Proxy.Get(cache)
	app.Service.Logger.Debug("Key: "+cache, " Cached value: "+cached.(string))

	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
	}
	fmt.Fprintf(writer, "%s", data)
}

// Request to get translation from Libretranslate service.
func (app *App) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {
	translate, _ := app.GetService().Translate(request.FormValue("q"), request.FormValue("source"), request.FormValue("target"))

	app.Service.Cached.Libre.Logger.Debug("POST request on localhost:8080/translate")
	fmt.Fprintf(writer, "%s", translate)
}

// Method to handle all requests
func (app *App) HandleRequests(port string) {
	//create a new router
	var (
		handler Handler = app
	)
	router := handler.GetRouter()
	app.Routes(router)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

// Routing,
//if you want to add new route, add it here
func (app *App) Routes(router *mux.Router) {
	var (
		handler Handler = app
	)
	router.HandleFunc("/languages", handler.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", handler.TranslatePageHandler).Methods("POST")
}
