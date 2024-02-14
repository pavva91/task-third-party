package api

import (
	"github.com/gorilla/mux"
)

var (
	Router *mux.Router
)

func NewRouter() {
	Router = mux.NewRouter()

	initializeRoutes()
}

func initializeRoutes() {
	task := Router.PathPrefix("/task").Subrouter()
	task.HandleFunc("", TasksHandler.CreateTask).Methods("POST")
	task.HandleFunc("/", TasksHandler.CreateTask).Methods("POST")
	task.HandleFunc("/{id:[0-9]+}", TasksHandler.GetByID).Methods("GET")
}
