package dto

import "errors"

type CreateTaskRequest struct {
	Method  string `json:"method"`
	// Location    string `json:"location"`
	URL  string `json:"url"`
	Headers    struct {
		Authentication string `json:"Authentication"`
	} `json:"headers"`
}

func (r *CreateTaskRequest) Validate() error {
	var errorMsg string

	if r.Method == "" {
		errorMsg = "Insert valid bucket name"
		err := errors.New(errorMsg)
		return err
	}

	// if r.Location == "" {
	// 	errorMsg = "Insert valid location"
	// 	err := errors.New(errorMsg)
	// 	return err
	// }

	if r.URL == "" {
		errorMsg = "Insert valid object name"
		err := errors.New(errorMsg)
		return err

	}

	return nil
}
