package main

import (
	"context"
	"fmt"
	"translateapp/internal"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	var dbstatus = internal.SetConnection()
	fmt.Println(dbstatus)
	log.Printf("starting...")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown")
}
