package dto

import (
	"github.com/pavva91/task-third-party/models"
)

type GetTaskResponse struct {
	ID             uint                   `json:"id"`
	Status         string                 `json:"status"`
	HttpStatusCode int                    `json:"httpStatusCode"`
	Headers        map[string]interface{} `json:"headers"` // TODO: headers from 3rd party service response
	Length         int                    `json:"length"`
}

func (dto *GetTaskResponse) ToDto(model models.Task) {
	dto.ID = model.ID
	dto.Status = model.Status.ToString()
	dto.HttpStatusCode = model.HttpStatusCode
	dto.Headers = model.ResHeaders
	dto.Length = model.Length
}
