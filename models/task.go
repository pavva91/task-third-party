package models

import (
	"github.com/pavva91/task-third-party/enums"
	"gorm.io/datatypes"
)

type Task struct {
	ID             uint
	URL            string
	Method         string
	HttpStatusCode int // <HTTP status of 3rd-party service response>
	// TODO: ReqHeaders data structure (array of key value pairs)
	ReqHeaders datatypes.JSONMap
	ResHeaders datatypes.JSONMap
	Status     enums.TaskStatus `gorm:"default:0"`
	Length     int
}
