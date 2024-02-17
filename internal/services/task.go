package services

import (
	"strconv"

	"github.com/pavva91/task-third-party/internal/models"
	"github.com/pavva91/task-third-party/internal/repositories"
)

var (
	Task Tasker = task{}
)

type Tasker interface {
	Create(task *models.Task) (*models.Task, error)
	GetByID(id uint) (*models.Task, error)
}

type task struct{}

func (s task) Create(task *models.Task) (*models.Task, error) {
	return repositories.Task.Create(task)
}

func (s task) GetByID(id uint) (*models.Task, error) {
	var task *models.Task
	strID := strconv.Itoa(int(id))
	task, err := repositories.Task.GetByID(strID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
