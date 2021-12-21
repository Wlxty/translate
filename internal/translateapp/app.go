package translateapp

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Application struct {
	Service Servicer
	Router  *mux.Router
	Logger  *zap.SugaredLogger
}

type Servicer interface {
	Languages() ([]Language, error)
	Translate(text, fromLang, toLang string) (Response, error)
}

func NewApp(service Servicer, logger *zap.SugaredLogger) *Application {
	router := mux.NewRouter().StrictSlash(true)

	return &Application{
		Service: service,
		Router:  router,
		Logger:  logger,
	}
}

func (s Application) LanguagePageHandler(writer http.ResponseWriter, request *http.Request) {
	langs, err := s.Service.Languages()
	if err != nil {
		// fmt.Fprintf(writer, string(jsonify))
		return
	}

	// Dump langs as JSON.
	// fmt.Fprintf(writer, )

	s.Logger.Debug("GET request on localhost:8080/languages")
}

func (s Application) TranslatePageHandler(writer http.ResponseWriter, request *http.Request) {
	text := request.FormValue("q")
	fromLang := request.FormValue("source")
	toLang := request.FormValue("target")

	response, err := s.Service.Translate(text, fromLang, toLang)
	if err != nil {
		// fmt.Fprintf(writer, string(jsonify))
		return
	}

	// Dump response as JSON.
	// fmt.Fprintf(writer, )

	s.Logger.Debug("POST request on localhost:8080/translate")

	// data := url.Values{
	// 	"q":      {request.FormValue("q")},
	// 	"source": {request.FormValue("source")},
	// 	"target": {request.FormValue("target")},
	// }
	// server.Service.Translate(writer, data)

	// logger := server.Logger
	// logger.Level = "Debug"
	// logger.Message(" POST request on localhost:8080/translate")
}

func (server *Application) HandleRequests(port string) {
	//create a new router
	router := server.Router

	Routes(router, server)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

func (server *Application) GetRouter() *mux.Router {
	return server.Router
}

func Routes(router *mux.Router, server *Application) error {
	router.HandleFunc("/languages", server.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", server.TranslatePageHandler).Methods("POST")
	return nil
}
