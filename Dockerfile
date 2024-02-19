FROM golang:1.21

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN go build -v -o /usr/local/bin/app ./...
RUN go build -v -o /usr/local/bin/app ./task-third-party/main/main.go

# Document that the service listens on port 8080.
EXPOSE 8080

CMD ["app"]
