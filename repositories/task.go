package repositories

import (
	"github.com/pavva91/task-third-party/db"
	"github.com/pavva91/task-third-party/models"
)

var (
	Task Tasker = task{}
)

type Tasker interface {
	Create(task *models.Task) (*models.Task, error)
	GetByID(id string) (*models.Task, error)
}

type task struct{}

func (r task) Create(task *models.Task) (*models.Task, error) {
	err := db.ORM.GetDB().Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r task) GetByID(id string) (*models.Task, error) {
	var task *models.Task
	err := db.ORM.GetDB().Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}
