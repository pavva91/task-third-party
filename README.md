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

### Run

1. Copy `./config/example-config.yml` into `./config/dev-config.yml` and put your own values

<a name="minio-docker"></a>

#### 2. Run Minio Docker Container

```bash
cd docker/minio
docker-compose up -d
```

3. Create `access-key-id` and `secret-access-key` from minio dashboard ([http://127.0.0.1:9001/access-keys/new-account](http://127.0.0.1:9001/access-keys/new-account)) and copy the values inside `./config/dev-config.yml`

4. Create `encryption-key-id` with:

```bash
kes key create dev-key
```

And copy the values inside `./config/dev-config.yml`

**_NOTE:_** To quickly install kes and mc I created a bash script in `./scripts/installMcAndKes.sh` ([run bash script](#bash))

**_NOTE:_** Guide to install and configure kes [install mc, kes and configure server side encryption](#kes)

#### Run Test Suite

1. With the [running minio docker container](#minio-docker) (verify with `docker ps`)

2. Create small and big files (inside project root folder):

```bash
mkdir testfiles
dd if=/dev/urandom of=./testfiles/verysmall1MiB bs=1M count=1
dd if=/dev/urandom of=./testfiles/small10MiB bs=1M count=10
dd if=/dev/urandom of=./testfiles/medium20MiB bs=1M count=20
dd if=/dev/urandom of=./testfiles/big100MiB bs=1M count=100
```

3. Run test suite (inside project root folder)

```bash
go test
```

#### Run Go Application

```bash
SERVER_ENVIRONMENT="dev" go run main.go
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
