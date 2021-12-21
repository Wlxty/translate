package main

import (
	"translateapp/internal/libretranslate"
	"translateapp/internal/logger"
	"translateapp/internal/translateapp"
)

func main() {
	//graceful shutdown to do
	logger := logger.NewLogger("debug", true)

	ltHost := "http://libretranslate:5000/"
	client := libretranslate.NewClient(logger, ltHost)
	service := translateapp.Service{
		logger,
		*client,
	}
	srv := translateapp.NewServer(&service)
	srv.HandleRequests(":8080")
}
