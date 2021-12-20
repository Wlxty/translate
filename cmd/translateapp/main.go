package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	server "translateapp/internal/server"
)

func main() {
	log.Printf("starting...")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	server.HandleRequests(":8080")
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown")

	//graceful shutdown to do
}
