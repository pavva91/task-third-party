package services

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pavva91/task-third-party/internal/enums"
	"github.com/pavva91/task-third-party/internal/models"
	"github.com/pavva91/task-third-party/internal/repositories"
)

func SendRequest(task *models.Task) (*models.Task, error) {

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
