package repositories

import (
	"github.com/pavva91/file-upload/db"
	"github.com/pavva91/file-upload/models"
)

var (
	Delegation DelegationRepositer = delegation{}
)

type DelegationRepositer interface {
	Create(task *models.Delegation) (*models.Delegation, error)
	List() ([]models.Delegation, error)
	GetByID(id string) (*models.Delegation, error)
}

type delegation struct{}

func (r delegation) Create(task *models.Delegation) (*models.Delegation, error) {
	err := db.ORM.GetDB().Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r delegation) List() ([]models.Delegation, error) {
	delegations := []models.Delegation{}
	err := db.ORM.GetDB().Order("timestamp DESC").Find(&delegations).Error
	if err != nil {
		return nil, err
	}
	return delegations, nil
}

func (r delegation) GetByID(id string) (*models.Delegation, error) {
	var task *models.Delegation
	err := db.ORM.GetDB().Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}
