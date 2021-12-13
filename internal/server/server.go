package server

import (
	"github.com/gorilla/mux"
	"translateapp/internal/service"
)

type Server struct {
	Service *service.Service
	Router  *mux.Router
}

func NewServer() *Server {

	router := mux.NewRouter().StrictSlash(true)
	var service service.Service
	return &Server{
		Service: &service,
		Router:  router,
	}
}
