package models

import (
	"github.com/pavva91/task-third-party/internal/enums"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model     `swaggerignore:"true"`
	URL            string
	Method         string
	HttpStatusCode int
	ReqHeaders     datatypes.JSONMap
	ResHeaders     datatypes.JSONMap
	Status         enums.TaskStatus `gorm:"default:0"`
	Length         int
}
