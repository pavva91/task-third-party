package services

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pavva91/task-third-party/enums"
	"github.com/pavva91/task-third-party/models"
	"github.com/pavva91/task-third-party/repositories"
)

var (
	Client Clienter = client{}
)

type Clienter interface {
	SendRequest(client *models.Task) (*models.Task, error)
}

type client struct{}

func (s client) SendRequest(task *models.Task) (*models.Task, error) {

	task.Status = enums.InProcess
	_, err := repositories.Task.UpdateTask(task)
	if err != nil {
		log.Println(err)
		return nil, err
	}

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
		return task, err
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
		task.HttpStatusCode = -1
		task.Length = -1
		repositories.Task.UpdateTask(task)

		log.Println(err)
		return task, err
	}
	defer res.Body.Close()

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
		return task, err
	}

	if res.StatusCode != 200 {
		task.Status = enums.Error
		task.HttpStatusCode = res.StatusCode
		task.Length = len(resBodyBytes)
		repositories.Task.UpdateTask(task)
		return task, err
	}

	task.Status = enums.Done
	task.HttpStatusCode = res.StatusCode
	task.Length = len(resBodyBytes)
	repositories.Task.UpdateTask(task)
	return task, nil
}
