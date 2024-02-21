package services

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pavva91/task-third-party/internal/enums"
	"github.com/pavva91/task-third-party/internal/models"
	"github.com/pavva91/task-third-party/internal/repositories"
)

var (
	Task Tasker = task{}
)

type Tasker interface {
	Create(task *models.Task) (*models.Task, error)
	GetByID(id uint) (*models.Task, error)
	SendRequest(task *models.Task) (*models.Task, error)
}

type task struct{}

func (s task) Create(task *models.Task) (*models.Task, error) {
	return repositories.Task.Create(task)
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

func (s task) SendRequest(task *models.Task) (*models.Task, error) {

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
		log.Println(err)

		task.Status = enums.Error
		task.HttpStatusCode = -1
		task.Length = -1

		_, dbErr := repositories.Task.UpdateTask(task)
		if dbErr != nil {
			log.Println(dbErr)
			return nil, dbErr
		}

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
		log.Println(err)

		task.Status = enums.Error
		task.HttpStatusCode = -1
		task.Length = -1

		_, dbErr := repositories.Task.UpdateTask(task)
		if dbErr != nil {
			log.Println(dbErr)
			return nil, dbErr
		}

		return task, err
	}
	defer res.Body.Close()

	log.Println("response status", res.Status)

	for k, v := range res.Header {
		task.ResHeaders[k] = v
	}

	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)

		task.Status = enums.Error
		task.HttpStatusCode = res.StatusCode
		task.Length = -1

		_, dbErr := repositories.Task.UpdateTask(task)
		if dbErr != nil {
			log.Println(dbErr)
			return nil, dbErr
		}

		return task, err
	}

	if res.StatusCode != 200 {
		log.Println(err)

		task.Status = enums.Error
		task.HttpStatusCode = res.StatusCode
		task.Length = len(resBodyBytes)

		_, dbErr := repositories.Task.UpdateTask(task)
		if dbErr != nil {
			log.Println(dbErr)
			return nil, dbErr
		}

		return task, err
	}

	task.Status = enums.Done
	task.HttpStatusCode = res.StatusCode
	task.Length = len(resBodyBytes)

	_, err = repositories.Task.UpdateTask(task)
	if err != nil {
		log.Println(err)
	}

	return task, nil
}
