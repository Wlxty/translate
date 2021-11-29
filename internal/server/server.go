package server
import(
	"translateapp/internal/languagesstore"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type taskServer struct {
	store *languagesstore.LanguagesStore
}

func NewTaskServer() *taskServer {
	store := languagesstore.New()
	return &taskServer{store: &store}
}

func HandleRequests() {
	server := NewTaskServer()

	//create a new router
	router := mux.NewRouter().StrictSlash(true)

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/languages", server.LanguagePageHandler).Methods("GET")
	router.HandleFunc("/", server.HomePageHandler).Methods("GET")
	router.HandleFunc("/translate", server.TranslatePageHandler).Methods("POST")
	//start and listen to requests
	log.Fatal(http.ListenAndServe(":8080", router))
}