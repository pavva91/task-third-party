package stubs

import "github.com/pavva91/task-third-party/internal/models"

type TaskRepository struct {
	CreateFn     func(*models.Task) (*models.Task, error)
	UpdateTaskFn func(*models.Task) (*models.Task, error)
	GetByIDFn    func() (*models.Task, error)
}

func (stub TaskRepository) Create(task *models.Task) (*models.Task, error) {
	return stub.CreateFn(task)
}

func (stub TaskRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	return stub.UpdateTaskFn(task)
}

func (stub TaskRepository) GetByID(id string) (*models.Task, error) {
	return stub.GetByIDFn()
}
