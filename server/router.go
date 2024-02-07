package server

import (
	"github.com/gorilla/mux"
	"github.com/pavva91/task-third-party/api"
)

var (
	r *mux.Router
)

func NewRouter() *mux.Router {
	r = mux.NewRouter()
	// TODO: CORS

	// NOTE: Routes
	r.HandleFunc("/", api.HomeHandler)

	task := r.PathPrefix("/task").Subrouter()
	task.HandleFunc("", api.TasksHandler.CreateTask).Methods("POST")
	task.HandleFunc("/", api.TasksHandler.CreateTask).Methods("POST")
	task.HandleFunc("/{id:[0-9]+}", api.TasksHandler.GetByID).Methods("GET")

	return r
}
