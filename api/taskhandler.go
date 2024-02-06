package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavva91/task-third-party/dto"
	"github.com/pavva91/task-third-party/errorhandlers"
	"github.com/pavva91/task-third-party/services"
)

type tasksHandler struct{}

var (
	TasksHandler = tasksHandler{}
)

func (h tasksHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Validation Request JSON
	err = body.Validate()
	if err != nil {
		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	// TODO: DTO to Model
	task := body.ToModel()

	// TODO: Create Task
	// TODO: Service
	services.Task.Create(task)
	// TODO: Send request to Third-Party (start goroutine)
	go services.Task.SendRequest(task)

	// TODO: Return Task ID
	var res dto.CreateTaskResponse
	res.ToDto(*task)

	js, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Write([]byte("task creation"))
	w.Write(js)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h tasksHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// TODO: Get Task by ID
	// w.Write([]byte("1234"))
	w.Write([]byte(id))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

