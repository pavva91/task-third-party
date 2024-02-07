// FileUploader.go MinIO example
package services

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pavva91/task-third-party/enums"
	"github.com/pavva91/task-third-party/models"
	"github.com/pavva91/task-third-party/repositories"
)

var (
	Task Tasker = task{}
)

type Tasker interface {
	Create(task *models.Task) (*models.Task, error)
	SendRequest(task *models.Task) (*models.Task, error)
	GetByID(id uint) (*models.Task, error)
}

type task struct{}

func (s task) Create(task *models.Task) (*models.Task, error) {
	return repositories.Task.Create(task)
}

func (s task) SendRequest(task *models.Task) (*models.Task, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest(task.Method, task.URL, nil)
	if err != nil {
		task.Status = enums.Error
		task.HttpStatusCode = -1
		task.Length = -1
		repositories.Task.UpdateTask(task)

		log.Println(err)
		return nil, err
	}
	// headers := datatypes.JSONQuery("headers")
	log.Println("Authorization value:", task.ReqHeaders)
	for k, v := range task.ReqHeaders {
		req.Header.Add(k, v.(string))
	}

	log.Println("header:", req.Header)

	res, err := client.Do(req)
	if err != nil {
		task.Status = enums.Error
		task.HttpStatusCode = res.StatusCode
		task.Length = -1
		repositories.Task.UpdateTask(task)

		log.Println(err)
		return nil, err
	}
	log.Println("response status", res.Status)

	for k, v := range res.Header {
		task.ResHeaders[k] = v
	}

	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		task.Status = enums.Error
		task.HttpStatusCode = res.StatusCode
		task.Length = -1
		repositories.Task.UpdateTask(task)

		log.Println(err)
		return nil, err
	}

	task.Length = len(resBodyBytes)
	task.HttpStatusCode = res.StatusCode
	task.Status = enums.Done
	repositories.Task.UpdateTask(task)
	return task, nil
}

func (s task) GetByID(id uint) (*models.Task, error) {
	var task *models.Task
	strID := strconv.Itoa(int(id))
	task, err := repositories.Task.GetByID(strID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
