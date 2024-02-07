# Programming Challenge: third party task

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
    "Authentication": "Basic
    bG9naW46cGFzc3dvcmQ=",
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

#### 1. Copy `./config/example-config.yml` into `./config/dev-config.yml` and put your own values

#### 2. Create .env for PostgreSQL credentials in `./docker/dev/.env` with this structure:

```bash
DB_USER=postgres
DB_PASSWORD=postgres
```

#### 3. Run PostgreSQL Docker Container

```bash
cd docker/dev
docker-compose up -d
```

#### 4. Run Go Application

```bash
SERVER_ENVIRONMENT="dev" go run main.go
```

#### Run Test Suite

##### Run with coverage

- Run all test suite and show coverage on browser: `go test ./... -coverprofile cover.out && go tool cover -html=cover.out`
- Run tests of ./dto and show coverage on browser: `go test -v -coverprofile cover.out ./dto/ && go tool cover -html=cover.out`
- Run test suite with "data race" detection : `go test --race ./...`

##### Run specific packages tests

- Run all test of the "api" package and show coverage on browser: `go test ./api -coverprofile cover.out && go tool cover -html=cover.out`
- Run all test of the "dto" package and show coverage on browser: `go test ./dto -coverprofile cover.out && go tool cover -html=cover.out`
- Run all test of the "services" package and show coverage on browser: `go test ./services -coverprofile cover.out && go tool cover -html=cover.out`
- Run all test of the "repositories" package and show coverage on browser: `go test ./repositories -coverprofile cover.out && go tool cover -html=cover.out`
- Run all test of the "utilities" package and show coverage on browser: `go test ./utilities -coverprofile cover.out && go tool cover -html=cover.out`

##### Run all test suite

```bash
go test
```

### cURL calls

#### Create Task

```bash
curl --location --request POST 'http://localhost:8080/files' \
--header 'Content-Type: application/json' \
--data-raw '{
    "bucketName":"test",
    "objectName": "big",
    "filepath": "./testfiles/big100MiB",
    "contentType": "application/octet-stream"
}'
```

#### Get Task status

```bash
curl --location --request GET 'http://localhost:8080/files' \
--header 'Content-Type: application/json' \
--data-raw '{
    "bucketName":"test",
    "objectName": "medium",
    "filepath": "./testfiles/medium20MiB",
    "contentType": "application/octet-stream"
}'
```
