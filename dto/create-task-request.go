package dto

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/pavva91/task-third-party/models"
	"gorm.io/datatypes"
)

type CreateTaskRequest struct {
	Method  string                 `json:"method"`
	URL     string                 `json:"url"`
	Headers map[string]interface{} `json:"headers"`
}

func (r *CreateTaskRequest) Validate() error {

	err := validateHttpHeaders(r)
	if err != nil {
		return err
	}

	err = validateURL(r)
	if err != nil {
		return err
	}

	err = validateHttpMethod(r)
	if err != nil {
		return err
	}

	return nil
}

func validateHttpHeaders(r *CreateTaskRequest) error {
	// NOTE: Header values are always a string

	for k, v := range r.Headers {
		switch t := v.(type) {
		case string:
			// do nothing
		default:
			err := errors.New(fmt.Sprintf("key %v doesn't have a string value, is of value %v", k, t))
			return err
		}

	}

	return nil
}

func validateHttpMethod(r *CreateTaskRequest) error {

	if r.Method == "" {
		err := errors.New("insert valid method name")
		return err
	}

	if r.Method != "GET" && r.Method != "POST" && r.Method != "PUT" && r.Method != "PATCH" && r.Method != "DELETE" {
		err := errors.New("not valid method name")
		return err
	}
	return nil
}

func validateURL(r *CreateTaskRequest) error {

	if r.URL == "" {
		err := errors.New("insert valid url")
		return err
	}

	u, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return err
	}

	if u.Scheme == "" {
		err := errors.New("insert url scheme")
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		err := errors.New("insert valid url scheme")
		return err
	}

	if u.Host == "" {
		err := errors.New("insert url host")
		return err
	}

	if (u.Host[0] >= 0 && u.Host[0] <= 47) || (u.Host[0] >= 58 && u.Host[0] <= 64) || (u.Host[0] >= 91 && u.Host[0] <= 96) || u.Host[0] >= 123 {
		log.Println("first:", u.Host[0])
		err := errors.New("insert url starting with char or number")
		return err
	}

	log.Println("URL:", u)

	return nil
}

func (dto *CreateTaskRequest) ToModel() *models.Task {
	var model models.Task
	model.ReqHeaders = datatypes.JSONMap(dto.Headers)
	model.URL = dto.URL
	model.Method = dto.Method
	return &model
}
