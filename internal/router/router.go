package router

import (
	"github.com/gorilla/mux"
	"github.com/pavva91/task-third-party/internal/handlers"
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
	task.HandleFunc("", handlers.TasksHandler.CreateTask).Methods("POST")
	task.HandleFunc("/", handlers.TasksHandler.CreateTask).Methods("POST")
	task.HandleFunc("/{id:[0-9]+}", handlers.TasksHandler.GetByID).Methods("GET")
}
