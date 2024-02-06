package models

import "github.com/pavva91/task-third-party/enums"

type Task struct {
	ID             uint
	HttpStatusCode uint // <HTTP status of 3rd-party service response>
	Headers        string // <headers array from 3rd-party service response>
	Status         enums.TaskStatus `gorm:"default:0"`
}
