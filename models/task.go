package models

import "github.com/pavva91/task-third-party/enums"

type Delegation struct {
	ID             uint
	Status enums.TaskStatus `gorm:"default:0"`
}
