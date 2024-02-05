package server

import (
	"github.com/gorilla/mux"
	"github.com/pavva91/file-upload/api"
)

var (
	r *mux.Router
)

func NewRouter() *mux.Router {
	r = mux.NewRouter()
	// TODO: CORS

	// NOTE: Routes
	r.HandleFunc("/", api.HomeHandler)

	s := r.PathPrefix("/task").Subrouter()
	s.HandleFunc("/", api.TasksHandler.CreateTask).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", api.TasksHandler.GetStatus).Methods("GET")

	return r
}
