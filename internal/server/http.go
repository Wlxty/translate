package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"translateapp/internal/dbcache"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logger"
	"translateapp/internal/translateapp"
)

//If you want to run api
// Running api in go routine and use of graceful shutdown
func Run() error {
	listenAddr := ":8080"

	logger := logger.NewLogger("debug", true)
	libre := libretranslate.NewClient(logger, "http://libretranslate:5000/")
	librewrapper := translateapp.NewLibreWrapper(libre)

	//DBCache
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//add constructor do ThroughDB
	cache := dbcache.NewThroughDB(conn, logger)
	cached := translateapp.NewCache(librewrapper, cache)
	service := translateapp.NewService(logger, cached)
	api := translateapp.NewApp(service)

	//cache := cache.Through{MemoryCache: cache.NewInMemoryCache()}
	//cached := translateapp.NewCache(librewrapper, &cache)
	//service := translateapp.NewService(logger, cached)
	//api := translateapp.NewApp(service)

	api.HandleRequests(":8080")
	server := http.Server{
		Addr:         listenAddr,
		Handler:      api,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main: API listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting server: %s", err)

	case <-shutdown:
		log.Println("main: Start shutdown")
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = server.Close()
		}
		if err != nil {
			return fmt.Errorf("main : could not stop server gracefully : %v", err)
		}
	}

	return nil
}
