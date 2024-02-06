// FileUploader.go MinIO example
package services

import (
	"github.com/pavva91/task-third-party/models"
	"github.com/pavva91/task-third-party/repositories"
)

var (
	Task Tasker = task{}
)

type Tasker interface {
	Create(task *models.Task) (*models.Task, error)
	SendRequest(task *models.Task) (*models.Task, error)
}

type task struct{}

func (s task) Create(task *models.Task) (*models.Task, error) {
	return repositories.Task.Create(task)
}

func (s task) SendRequest(task *models.Task) (*models.Task, error) {
	// TODO: send request to third-party service
	// TODO: create http client
	// TODO: update task state accordingly
	// TODO: send request to third-party service
	// TODO: wait for response (this function is run async)
	// TODO: update task state accordingly
	return task, nil
}
