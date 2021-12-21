package main

import (
	"fmt"
	"log"
	"os"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
)

func main() {
	// os.Getenv("TRANSLATEAPP_LOGLEVE")
	// os.Getenv("TRANSLATEAPP_ENABLE_DEVEL_LOGS")
	logger := logging.NewLogger("debug", true)

	ltHost := os.Getenv("TRANSLATEAPP_LTHOST") // "http://libretranslate:5000"

	client := libretranslate.NewClient(ltHost, logger)

	service := translateapp.NewService(client, logger)

	app := translateapp.NewApp(service, logger)

	srv := server.New(app, logger)

	// srv := &http.Server{
	// 	Handler: s.GetRouter(),
	// 	Addr:    "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// go func() {
	// 	s.HandleRequests(":8080")
	// }()

	// log.Println("server started")

	// stopC := make(chan os.Signal)
	// signal.Notify(stopC, os.Interrupt)
	// <-stopC

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// log.Println("server stopping...")
	// defer cancel()

	// log.Fatal(srv.Shutdown(ctx))

	port := os.Getenv("TRANSLATEAPP_PORT")

	log.Fatal(srv.Run(fmt.Sprintf(":%d", port)))
}
