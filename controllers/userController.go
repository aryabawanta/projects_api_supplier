package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) MakeRoutes() {
	var class_url = "/users"

	routes := s.Router.PathPrefix("/v1").Subrouter()

	routes.HandleFunc(class_url, list).Methods("GET")

	return routes
}

func list(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "list"}`))
}
