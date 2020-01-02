package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize(){
	server.Router = mux.NewRouter()

	server.Router.PathPrefix("/api").Subrouter()
	
	server.MakeRoutes()
}