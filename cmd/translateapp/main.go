package main

import (
	"translateapp/internal/server"
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("starting...")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	server.HandleRequests()
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown")
}
