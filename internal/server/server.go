package server
import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func HandleRequests(port string) {
	server := NewServer()

	//create a new router
	router := server.router
	Routes(router, server)
	//start and listen to requests
	log.Fatal(http.ListenAndServe(port, router))
}

func Routes(router *mux.Router, server *Server) (error){
        router.HandleFunc("/languages", server.LanguagePageHandler).Methods("GET")
        router.HandleFunc("/translate", server.TranslatePageHandler).Methods("POST")
        return nil
}

