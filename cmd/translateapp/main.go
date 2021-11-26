package main

import (
	//"context"
	"fmt"
	"translateapp/internal"
	"log"
	//"os/signal"
	//"syscall"
	"net/http"
	servers "translateapp/internal/server"
	"github.com/gorilla/mux"
)

func main() {
	var dbstatus = internal.SetConnection()
	fmt.Println(dbstatus)
	log.Printf("starting...")
	server := servers.NewTaskServer()

	//create a new router
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/languages/", server.TaskHandler).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

	/* ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown") */
}
