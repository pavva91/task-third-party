package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pavva91/task-third-party/dto"
	"github.com/pavva91/task-third-party/errorhandlers"
	"github.com/pavva91/task-third-party/services"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type tasksHandler struct{}

var (
	TasksHandler = tasksHandler{}
)

func (h tasksHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		if err.Error() == "EOF" {
			err = errors.New("insert valid json body")
		}
		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	err = body.Validate()
	if err != nil {
		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	task := body.ToModel()
	task.ResHeaders = datatypes.JSONMap(make(map[string]interface{}))

	task, err = services.Task.Create(task)
	if err != nil {
		log.Println(err)
		errorhandlers.InternalServerErrorHandler(w, r)
		return
	}

	go services.Client.SendRequest(task)

	var res dto.CreateTaskResponse
	res.ToDto(*task)

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	w.Header().Set("Content-Type", "application/json")
}

func (h tasksHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	strID := mux.Vars(r)["id"]

	i, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		errorhandlers.BadRequestHandler(w, r, errors.New("insert valid id"))
		return
	}
	id := uint(i)

	task, err := services.Task.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorhandlers.NotFoundHandler(w, r, err)
			return
		}

		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	var res dto.GetTaskResponse
	res.ToDto(*task)

	js, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		errorhandlers.InternalServerErrorHandler(w, r)
		return
	}

	w.Write(js)
	w.Header().Set("Content-Type", "application/json")
}
