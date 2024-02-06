package dto

import (
	"github.com/pavva91/task-third-party/models"
)

type CreateTaskResponse struct {
	ID uint `json:"id"`
}

func (dto *CreateTaskResponse) ToDto(userModel models.Task) {
	dto.ID = userModel.ID
}
