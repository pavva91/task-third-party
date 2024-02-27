# Programming Challenge: third party task

[![Go](https://github.com/pavva91/task-third-party/actions/workflows/go.yml/badge.svg)](https://github.com/pavva91/task-third-party/actions/workflows/go.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/pavva91/task-third-party)](https://goreportcard.com/report/github.com/pavva91/task-third-party)

## Description

Write HTTP server for a service that would make http requests to 3rd-party
services.
Work algorithm.
The client sends a task to the service to perform an http request to a 3rd-party
services. The task is described in json format, the generated task id is returned
in response and its execution starts in the background.

Request example to service
Request:

- POST /task

```json
{
  "method": "GET",
  "url": "http://google.com",
  "headers": {
    "Authentication": "Basic bG9naW46cGFzc3dvcmQ="
  }
}
```

Response:
200 OK

```json
{
  "id": <generated unique id>
}
```

The client must have a method that can be used to find out the status of the
task.
Request example
Request:

- GET task/<taskId>

Response:

200 OK

```json
{
    "id": <unique id>,
    "status": "done/in_process/error/new"
    "httpStatusCode": <HTTP status of 3rd-party service response>,
    "headers": {
        <headers array from 3rd-party service response>
    },
    "length": <content length of 3rd-party service response>
}
```

We'd like to see code close to production with clear variable names and http
routes, unit tests, etc.

## Solution

I use PostgreSQL to persist data.

### Run

#### 1. Copy `./example.env` into `./.env` and put your own values:

```bash
DB_USER=postgres
DB_PASSWORD=postgres
```

#### 2. Build go docker image

Build [Dockerfile](./Dockerfile)

on root of the project run:

```bash
docker compose build
```

#### 3. Run everything

Run [docker-compose.yml](./docker-compose.yml)

on root of the project run:

```bash
docker compose up
```

#### Run Test Suite

##### Run with coverage

- Run all test suite and show coverage on browser: `go test ./... -coverprofile cover.out && go tool cover -html=cover.out`
- Run tests of ./dto and show coverage on browser: `go test -v -coverprofile cover.out ./internal/dto/ && go tool cover -html=cover.out`
- Run all test of the "dto" package and show coverage on browser: `go test ./internal/dto -coverprofile cover.out && go tool cover -html=cover.out`
- Run test suite with "data race" detection : `go test --race ./...`

##### Run all test suite

```bash
go test ./...
```

### Run Development Environment

#### 1. Copy `./config/example-config.yml` into `./config/dev-config.yml` and put your own values

#### 2. Create .env for PostgreSQL credentials in `./docker/dev/.env` with this structure:

```bash
DB_USER=postgres
DB_PASSWORD=postgres
```

#### 3. Run Docker Compose

Runs PostgreSQL and Go application with air hot-reload.

```bash
docker-compose -f docker/dev/docker-compose.yml up
```

### Swagger API

A Swagger API is available at:

- [http://localhost:8080/swagger/index.html#/](http://localhost:8080/swagger/index.html#/)

### cURL calls

#### Create Task

```bash
curl --location --request POST 'http://localhost:8080/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "method": "GET",
    "url": "https://example.com",
    "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ="
    }
}'
```

#### Get Task status

```bash
curl --location --request GET 'http://localhost:8080/task/1' \
--header 'Content-Type: application/json'
```

### Run linter

```bash
staticcheck ./...
```

#### Metalinter

Useful meta-linter:

- [https://golangci-lint.run/](https://golangci-lint.run/)

##### Install

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.56.1
```

##### Run

```bash
golangci-lint run
```

### Makefile

Source: [https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects](https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects)
Gist: [https://gist.github.com/alexedwards/3b40775846535d0014ab1ff477e4a568](https://gist.github.com/alexedwards/3b40775846535d0014ab1ff477e4a568)

```bash
SERVER_ENVIRONMENT="dev" make run
```

### Dockerfile

Source 1: [https://docs.docker.com/develop/develop-images/dockerfile_best-practices/](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
Source 2: [https://hub.docker.com/\_/golang/](https://hub.docker.com/_/golang/)

#### Build image

```bash
docker build -t my-golang-app .
```

#### Run docker container

```bash
docker run -it --rm --name my-running-app -e "SERVER_ENVIRONMENT=dev" -p 6060:8080 my-golang-app
```

### Docker Compose

Run both db and go binary with a single docker-compose file:
From the root of the project run:

```bash
docker compose up
```

### Systemd Daemon Unit

#### Create symlink

```bash
sudo ln -s ~/work/task/daemons/gotask.service /etc/systemd/system/gotask.service
```

#### Start Service

```bash
sudo systemctl start gotask.service
```

## Kubernetes

Use kompose to convert docker-compose into k8s configurations.
From project root do:

```bash
kompose covert /folder/to/store/k8s/configs
```

# Deploy to kubernetes

```bash
kubectl apply -f /folder/to/store/k8s/configs
```
