package dto

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/pavva91/task-third-party/internal/models"
	"gorm.io/datatypes"
)

type CreateTaskRequest struct {
	Method  string                 `json:"method"`
	URL     string                 `json:"url"`
	Headers map[string]interface{} `json:"headers"`
}

func (r *CreateTaskRequest) Validate() error {
	// FIX: Validation per character or number looks strange. Could have been done by regular expression or standard functions from unicode library.

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
			err := fmt.Errorf("key %v doesn't have a string value, is of value %v", k, t)
			return err
		}

	}

	return nil
}

func validateHttpMethod(req *CreateTaskRequest) error {

	if req.Method == "" {
		err := errors.New("insert valid method name")
		return err
	}

	if req.Method != "GET" && req.Method != "POST" && req.Method != "PUT" && req.Method != "PATCH" && req.Method != "DELETE" {
		err := errors.New("not valid method name")
		return err
	}
	return nil
}

func validateURL(req *CreateTaskRequest) error {

	if req.URL == "" {
		err := errors.New("insert valid url")
		return err
	}

	u, err := url.ParseRequestURI(req.URL)
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

	// NOTE: validation with regex library
	var validURL = regexp.MustCompile(`^[a-zA-Z0-9]`)
	if !validURL.MatchString(u.Host) {
		log.Println("first:", u.Host[0])
		err := errors.New("insert url starting with char or number")
		return err
	}
	log.Println(validURL)

	// NOTE: validation with unicode library
	// var (
	// 	hasNum bool
	// 	hasUp  bool
	// 	hasLow bool
	// )
	// for _, r := range u.Host[0:1] {
	// 	switch {
	// 	case unicode.IsDigit(r):
	// 		hasNum = true
	// 	case unicode.IsUpper(r):
	// 		hasUp = true
	// 	case unicode.IsLower(r):
	// 		hasLow = true
	// 	}
	// }
	// if !hasNum && !hasUp && !hasLow {
	// 	log.Println("first:", u.Host[0])
	// 	err := errors.New("insert url starting with char or number")
	// 	return err
	// }

	// NOTE: was correct but better using regex or unicode libraries as above
	// if (u.Host[0] <= 47) || (u.Host[0] >= 58 && u.Host[0] <= 64) || (u.Host[0] >= 91 && u.Host[0] <= 96) || u.Host[0] >= 123 {
	// 	log.Println("first:", u.Host[0])
	// 	err := errors.New("insert url starting with char or number")
	// 	return err
	// }

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
