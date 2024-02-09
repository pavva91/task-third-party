package stubs

import "github.com/pavva91/task-third-party/models"

type TaskService struct {
	CreateFn     func(*models.Task) (*models.Task, error)
	GetByIDFn    func(uint) (*models.Task, error)
}

func (stub TaskService) Create(task *models.Task) (*models.Task, error) {
	return stub.CreateFn(task)
}

func (stub TaskService) GetByID(id uint) (*models.Task, error) {
	return stub.GetByIDFn(id)
}
