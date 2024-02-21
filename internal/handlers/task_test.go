package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pavva91/task-third-party/internal/dto"
	"github.com/pavva91/task-third-party/internal/enums"
	"github.com/pavva91/task-third-party/internal/models"
	"github.com/pavva91/task-third-party/internal/services"
	"github.com/pavva91/task-third-party/internal/stubs"
	"gorm.io/datatypes"
	"gorm.io/gorm/logger"
)

func Test_tasksHandler_GetByID(t *testing.T) {
	taskStub := &models.Task{
		URL:        "http://foo.bar",
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
		Status: enums.New,
	}

	stubFail := stubs.TaskService{}
	stubFail.GetByIDFn = func(uint) (*models.Task, error) {
		return nil, errors.New("stub error")
	}

	stubNotFound := stubs.TaskService{}
	stubNotFound.GetByIDFn = func(uint) (*models.Task, error) {
		return nil, logger.ErrRecordNotFound
	}

	stubOk := stubs.TaskService{}
	stubOk.GetByIDFn = func(uint) (*models.Task, error) {
		return taskStub, nil
	}

	type args struct {
		vars map[string]string
	}
	tests := map[string]struct {
		args             args
		stub             stubs.TaskService
		wantErr          bool
		expectedHttpCode int
		expectedResBody  string
	}{
		"without id": {
			args{
				vars: map[string]string{
					"id": "",
				},
			},
			stubs.TaskService{},
			true,
			400,
			"insert valid id",
		},
		"not integer id": {
			args{
				vars: map[string]string{
					"id": "abcd",
				},
			},
			stubs.TaskService{},
			true,
			400,
			"insert valid id",
		},
		"get by id fail": {
			args{
				vars: map[string]string{
					"id": "1",
				},
			},
			stubFail,
			true,
			400,
			"stub error",
		},
		"task not found": {
			args{
				vars: map[string]string{
					"id": "1",
				},
			},
			stubNotFound,
			true,
			404,
			"record not found",
		},
		"ok": {
			args{
				vars: map[string]string{
					"id": "1",
				},
			},
			stubOk,
			false,
			200,
			fmt.Sprintf(`{"id":%d,"status":"new","httpStatusCode":%d,"headers":{},"length":%d}`, taskStub.ID, taskStub.HttpStatusCode, taskStub.Length),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			services.Task = tt.stub

			r, err := http.NewRequest("GET", "/task", nil)
			if err != nil {
				t.Fatal(err)
			}
			r = mux.SetURLVars(r, tt.args.vars)

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(TasksHandler.GetByID)
			handler.ServeHTTP(w, r)

			if (w.Code != 200) != tt.wantErr {
				t.Errorf("CreateTaskRequest.Validate() error = %v, wantErr %v", w.Code, tt.wantErr)
			}
			if w.Code != tt.expectedHttpCode {
				t.Errorf("CreateTaskRequest.Validate() error = %v, wantErr %v", w.Code, tt.expectedHttpCode)
			}
			if w.Body.String() != tt.expectedResBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					w.Body.String(), tt.expectedResBody)
			}
		})
	}
}

func Test_tasksHandler_CreateTask(t *testing.T) {

	taskStub := &models.Task{
		URL:        "http://foo.bar",
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
		Status: enums.New,
	}

	stubFail := stubs.TaskService{}
	stubFail.CreateFn = func(*models.Task) (*models.Task, error) {
		return nil, errors.New("error")
	}

	stubFailReq := stubs.TaskService{}
	stubFailReq.CreateFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	stubFailReq.SendRequestFn = func(*models.Task) (*models.Task, error) {
		return nil, errors.New("request error")
	}

	stubOk := stubs.TaskService{}
	stubOk.CreateFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	stubOk.SendRequestFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}

	type args struct {
		vars    map[string]string
		reqBody interface{}
	}
	tests := map[string]struct {
		args             args
		stub             stubs.TaskService
		wantErr          bool
		expectedHttpCode int
		expectedResBody  string
	}{
		"failed validation": {
			args{
				vars:    map[string]string{},
				reqBody: dto.CreateTaskRequest{},
			},
			stubs.TaskService{},
			true,
			400,
			"insert valid url",
		},
		"create task return error": {
			args{
				vars: map[string]string{},
				reqBody: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://example",
				},
			},
			stubFail,
			true,
			500,
			"500 Internal Server Error",
		},
		"third paryt request error": {
			args{
				vars: map[string]string{},
				reqBody: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://example",
				},
			},
			stubFailReq,
			false,
			200,
			fmt.Sprintf(`{"id":%d}`, taskStub.ID),
		},
		"ok": {
			args{
				vars: map[string]string{},
				reqBody: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://example",
				},
			},
			stubOk,
			false,
			200,
			fmt.Sprintf(`{"id":%d}`, taskStub.ID),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			services.Task = tt.stub

			marshalled, err := json.Marshal(tt.args.reqBody)
			if err != nil {
				log.Fatalf("impossible to marshall task: %s", err)
			}
			log.Print(marshalled)

			r, err := http.NewRequest("POST", "/task", bytes.NewReader(marshalled))
			if err != nil {
				t.Fatal(err)
			}
			r = mux.SetURLVars(r, tt.args.vars)
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(TasksHandler.CreateTask)
			handler.ServeHTTP(w, r)

			if (w.Code != 200) != tt.wantErr {
				t.Errorf("CreateTaskRequest.Validate() error = %v, wantErr %v", w.Code, tt.wantErr)
			}
			if w.Code != tt.expectedHttpCode {
				t.Errorf("CreateTaskRequest.Validate() error = %v, wantErr %v", w.Code, tt.expectedHttpCode)
			}
			if w.Body.String() != tt.expectedResBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					w.Body.String(), tt.expectedResBody)
			}
		})
	}
}
