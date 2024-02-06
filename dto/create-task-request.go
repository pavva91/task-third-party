package dto

import (
	"errors"

	"github.com/pavva91/task-third-party/models"
)

type CreateTaskRequest struct {
	Method string `json:"method"`
	// Location    string `json:"location"`
	URL     string `json:"url"`
	Headers string `json:"headers"`
	// Headers struct {
	// 	Authentication string `json:"Authentication"`
	// } `json:"headers"`
}

func (r *CreateTaskRequest) Validate() error {

	err := validateHttpMethod(r)
	if err != nil {
		return err
	}

	if r.URL == "" {
		err := errors.New("insert valid url")
		return err

	}

	return nil
}

func validateHttpMethod(r *CreateTaskRequest) error {
	if r.Method == "" {
		err := errors.New("insert valid method name")
		return err
	}

	if r.Method != "GET" && r.Method != "POST" && r.Method != "PUT" && r.Method != "PATCH" {
		err := errors.New("not valid method name")
		return err
	}
	return nil
}

func (dto *CreateTaskRequest) ToModel() *models.Task {
	var model models.Task
	model.Headers = dto.Headers
	// TODO: Assign values to model
	return &model
}
