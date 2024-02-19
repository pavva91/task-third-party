package services

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pavva91/task-third-party/internal/enums"
	"github.com/pavva91/task-third-party/internal/models"
	"github.com/pavva91/task-third-party/internal/repositories"
	"github.com/pavva91/task-third-party/internal/stubs"
	"gorm.io/datatypes"
)

func Test_SendRequest_Error_UpdateTaskError(t *testing.T) {
	expectedError := "stub error"

	taskStub := &models.Task{
		URL:        "",
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return nil, errors.New("stub error")
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	test := struct {
		args     args
		want     *models.Task
		wantErr  bool
		errorMsg string
	}{
		args{
			taskStub,
		},
		nil,
		true,
		expectedError,
	}
	t.Run("error updating record db", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got != test.want || got != nil {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if err != nil {
			if err.Error() != test.errorMsg {
				t.Errorf("SendRequest() = %v, want %v", err.Error(), test.errorMsg)
			}
		}
	})
}

func Test_SendRequest_Error_NoSchemaURL(t *testing.T) {
	wrongURL := "/no.schema"
	expectedError := fmt.Sprintf("Get \"%v\": unsupported protocol scheme \"\"", wrongURL)

	taskStub := &models.Task{
		URL:        wrongURL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	expected := &models.Task{
		URL: wrongURL,
		ResHeaders: datatypes.JSONMap(map[string]interface{}{
			"Content-Length": "[0]",
			"Date":           "[Thu, 08 Feb 2024 00:00:50 GMT]",
		}),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
		HttpStatusCode: -1,
		Length:         -1,
		Status:         enums.Error,
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	test := struct {
		args     args
		want     *models.Task
		wantErr  bool
		errorMsg string
	}{
		args{
			taskStub,
		},
		expected,
		true,
		expectedError,
	}
	t.Run("no schema url", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Status != test.want.Status {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if err != nil {
			if err.Error() != test.errorMsg {
				t.Errorf("SendRequest() = %v, want %v", err.Error(), test.errorMsg)
			}
		}
	})
}

func Test_SendRequest_Error_WrongSchemaURL(t *testing.T) {
	wrongURL := "wrongschema://example.com"
	expectedError := fmt.Sprintf("Get \"%v\": unsupported protocol scheme \"%v\"", wrongURL, "wrongschema")

	taskStub := &models.Task{
		URL:        wrongURL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	expected := &models.Task{
		URL: wrongURL,
		ResHeaders: datatypes.JSONMap(map[string]interface{}{
			"Content-Length": "[0]",
			"Date":           "[Thu, 08 Feb 2024 00:00:50 GMT]",
		}),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
		HttpStatusCode: -1,
		Length:         -1,
		Status:         enums.Error,
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	test := struct {
		args     args
		want     *models.Task
		wantErr  bool
		errorMsg string
	}{
		args{
			taskStub,
		},
		expected,
		true,
		expectedError,
	}
	t.Run("wrong schema url", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Status != test.want.Status {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if err != nil {
			if err.Error() != test.errorMsg {
				t.Errorf("SendRequest() = %v, want %v", err.Error(), test.errorMsg)
			}
		}
	})
}

func Test_SendRequest_Error_URLMissingSlash(t *testing.T) {
	wrongURL := "http:/example.org"
	expectedError := fmt.Sprintf("Get \"%v\": http: no Host in request URL", wrongURL)

	taskStub := &models.Task{
		URL:        wrongURL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	expected := &models.Task{
		URL: wrongURL,
		ResHeaders: datatypes.JSONMap(map[string]interface{}{
			"Content-Length": "[0]",
			"Date":           "[Thu, 08 Feb 2024 00:00:50 GMT]",
		}),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
		HttpStatusCode: -1,
		Length:         -1,
		Status:         enums.Error,
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	test := struct {
		args     args
		want     *models.Task
		wantErr  bool
		errorMsg string
	}{
		args{
			taskStub,
		},
		expected,
		true,
		expectedError,
	}
	t.Run("missing slash url", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Status != test.want.Status {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if err != nil {
			if err.Error() != test.errorMsg {
				t.Errorf("SendRequest() = %v, want %v", err.Error(), test.errorMsg)
			}
		}
	})
}

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
		URL:        srv.URL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	expected := &models.Task{
		URL: srv.URL,
		ResHeaders: datatypes.JSONMap(map[string]interface{}{
			"Content-Length": "[0]",
			"Date":           "[Thu, 08 Feb 2024 00:00:50 GMT]",
		}),
		HttpStatusCode: 500,
		Status:         enums.Error,
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	test := struct {
		server  *httptest.Server
		args    args
		want    *models.Task
		wantErr bool
	}{
		srv,
		args{
			taskStub,
		},
		expected,
		false,
	}
	t.Run("3rd returns 500", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		// if !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
		// }
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Status != test.want.Status {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.HttpStatusCode != test.want.HttpStatusCode {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.URL != test.want.URL {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		// if got.ResHeaders["Content-Length"] != tt.want.ResHeaders["Content-Length"] {
		// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
		// }
	})
}

func Test_SendRequest_ThirdPartyDown(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("Expected to request '/', got: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusInternalServerError)
	}))
	srv.Close()

	taskStub := &models.Task{
		URL:        srv.URL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	expected := &models.Task{
		URL: srv.URL,
		ResHeaders: datatypes.JSONMap(map[string]interface{}{
			"Content-Length": "[0]",
			"Date":           "[Thu, 08 Feb 2024 00:00:50 GMT]",
		}),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
		HttpStatusCode: -1,
		Length:         -1,
		Status:         enums.Error,
	}

	expectedError := fmt.Sprintf(`Get "%v": dial tcp %v: connect: connection refused`, expected.URL, expected.URL[7:])

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
	test := struct {
		server   *httptest.Server
		args     args
		want     *models.Task
		wantErr  bool
		errorMsg string
	}{
		srv,
		args{
			taskStub,
		},
		expected,
		true,
		expectedError,
	}
	t.Run("3rd returns 500", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		// if !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
		// }
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Status != test.want.Status {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.HttpStatusCode != test.want.HttpStatusCode {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.URL != test.want.URL {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if err != nil {
			if err.Error() != test.errorMsg {
				t.Errorf("SendRequest() = %v, want %v", err.Error(), test.errorMsg)
			}
		}
		// if got.ResHeaders["Content-Length"] != tt.want.ResHeaders["Content-Length"] {
		// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
		// }
	})
}

func Test_SendRequest_OK200(t *testing.T) {
	resBody := `{"foo":"a"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("Expected to request '/', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, resBody)
		// w.Write(unexpectedJSON)
	}))
	defer srv.Close()

	taskStub := &models.Task{
		URL:        srv.URL,
		ResHeaders: datatypes.JSONMap(make(map[string]interface{})),
		ReqHeaders: datatypes.JSONMap(map[string]interface{}{
			"foo": "alpha",
			"bar": "beta",
		}),
	}

	expected := &models.Task{
		URL: srv.URL,
		ResHeaders: datatypes.JSONMap(map[string]interface{}{
			"Content-Length": "[0]",
			"Date":           "[Thu, 08 Feb 2024 00:00:50 GMT]",
		}),
		HttpStatusCode: 200,
		Status:         enums.Done,
		Length:         len(resBody) + 1,
	}

	taskRepositoryStub := stubs.TaskRepository{}
	taskRepositoryStub.UpdateTaskFn = func(*models.Task) (*models.Task, error) {
		return taskStub, nil
	}
	repositories.Task = taskRepositoryStub

	type args struct {
		task *models.Task
	}
	test := struct {
		server  *httptest.Server
		args    args
		want    *models.Task
		wantErr bool
	}{
		srv,
		args{
			taskStub,
		},
		expected,
		false,
	}
	t.Run("3rd returns 200", func(t *testing.T) {
		got, err := SendRequest(test.args.task)
		if (err != nil) != test.wantErr {
			t.Errorf("SendRequest() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		// if !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
		// }
		if got.ID != test.want.ID {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Status != test.want.Status {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.HttpStatusCode != test.want.HttpStatusCode {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.URL != test.want.URL {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		if got.Length != test.want.Length {
			t.Errorf("SendRequest() = %v, want %v", got, test.want)
		}
		// if got.ResHeaders["Content-Length"] != tt.want.ResHeaders["Content-Length"] {
		// 	t.Errorf("SendRequest() = %v, want %v", got, tt.want)
		// }
	})
}
