package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pavva91/task-third-party/enums"
	"github.com/pavva91/task-third-party/models"
	"github.com/pavva91/task-third-party/repositories"
	"github.com/pavva91/task-third-party/stubs"
	"gorm.io/datatypes"
)

func Test_SendRequest_Error500(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("Expected to request '/', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusInternalServerError)
		// w.Write(unexpectedJSON)
	}))
	defer srv.Close()

	taskStub := &models.Task{
		ID:         1,
		URL:        srv.URL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo":"alpha",
			"bar":"beta",
		}),
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.CreateFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	tests := map[string]struct {
		server  *httptest.Server
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		"3rd internal server error": {
			srv,
			args{
				taskStub,
			},
			&models.Task{
				ID:         1,
				URL:        srv.URL,
				ResHeaders: datatypes.JSONMap(map[string]interface{}{
					"Content-Length":"[0]",
					"Date":"[Thu, 08 Feb 2024 00:00:50 GMT]",
				}),
				// ResHeaders: datatypes.JSONMap(map[string]interface{}{
				// 	"Content-Length":"[0]",
				// 	"Date":"[Thu, 08 Feb 2024 00:00:50 GMT]",
				// }),
				HttpStatusCode: 500,
				Status: enums.Done,
			},
			false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := SendRequest(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			// }
			if got.ID != tt.want.ID {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			if got.Status != tt.want.Status {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			if got.HttpStatusCode != tt.want.HttpStatusCode {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			if got.URL != tt.want.URL {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			// if got.ResHeaders["Content-Length"] != tt.want.ResHeaders["Content-Length"] {
			// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_SendRequest_OK200(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("Expected to request '/', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		// w.Write(unexpectedJSON)
	}))
	defer srv.Close()

	taskStub := &models.Task{
		ID:         1,
		URL:        srv.URL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.CreateFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	tests := map[string]struct {
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		"ok": {
			args{
				taskStub,
			},
			&models.Task{
				ID:         1,
				URL:        srv.URL,
				ResHeaders: datatypes.JSONMap(map[string]interface{}{
					"Content-Length":"[0]",
					"Date":"[Thu, 08 Feb 2024 00:00:50 GMT]",
				}),
				// ResHeaders: datatypes.JSONMap(map[string]interface{}{
				// 	"Content-Length":"[0]",
				// 	"Date":"[Thu, 08 Feb 2024 00:00:50 GMT]",
				// }),
				HttpStatusCode: 200,
				Status: enums.Done,
			},
			false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := SendRequest(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			// }
			if got.ID != tt.want.ID {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			if got.Status != tt.want.Status {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			if got.HttpStatusCode != tt.want.HttpStatusCode {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			if got.URL != tt.want.URL {
				t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			}
			// if got.ResHeaders["Content-Length"] != tt.want.ResHeaders["Content-Length"] {
			// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
			// }
		})
	}
}
