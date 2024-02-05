package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavva91/task-third-party/dto"
	"github.com/pavva91/task-third-party/errorhandlers"
)

type tasksHandler struct{}

var (
	TasksHandler = tasksHandler{}
)

// var (
// 	TaskRe         = regexp.MustCompile(`^/task/*$`)
// 	TaskReWithID   = regexp.MustCompile(`^/task/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
// 	TaskReWithName = regexp.MustCompile(`^/task/.+$`)
// )

func (h tasksHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqBody dto.CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = reqBody.Validate()
	if err != nil {
		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	// TODO: Create Task
	// TODO: Send request to Third-Party (start goroutine)
	// TODO: Return Task ID

	w.Write([]byte("task creation"))
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

