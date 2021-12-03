package server
import(
	"log"
	"net/http"
)

func HandleRequests() {
	server := NewServer()

	//create a new router
	router := server.router
	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/languages", server.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/translate", server.TranslatePageHandler).Methods("POST")
	//start and listen to requests
	log.Fatal(http.ListenAndServe(":8080", router))
}
