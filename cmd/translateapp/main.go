package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"translateapp/internal/server"
)

func main() {
	//graceful shutdown to do
	s := server.NewServer()
	srv := &http.Server{
		Handler: s.GetRouter(),
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		s.HandleRequests(":8080")
	}()
	log.Println("server started")

	stopC := make(chan os.Signal)
	signal.Notify(stopC, os.Interrupt)
	<-stopC

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("server stopping...")
	defer cancel()

	log.Fatal(srv.Shutdown(ctx))
}
