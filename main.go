package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/pavva91/task-third-party/api"
	"github.com/pavva91/task-third-party/config"
	"github.com/pavva91/task-third-party/db"
	"github.com/pavva91/task-third-party/models"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	isDebug := false
	if len(os.Args) == 2 {
		debugArg := os.Args[1]
		if debugArg == "d" || debugArg == "debug" {
			os.Setenv("SERVER_ENVIRONMENT", "dev")
			isDebug = true
		}
	}
	log.Println(fmt.Sprintf("debug mode: %t", isDebug))

	api.NewRouter()

	env := os.Getenv("SERVER_ENVIRONMENT")

	log.Println(fmt.Sprintf("Running Environment: %s", env))

	switch env {
	case "dev":
		setConfig("./config/dev-config.yml")
	case "stage":
		log.Panic(fmt.Sprintf("Incorrect Dev Environment: %s\nInterrupt execution", env))
	case "prod":
		log.Panic(fmt.Sprintf("Incorrect Dev Environment: %s\nInterrupt execution", env))
	default:
		log.Panic(fmt.Sprintf("Incorrect Dev Environment: %s\nInterrupt execution", env))
	}

	// connect to db
	db.ORM.MustConnectToDB(config.ServerConfigValues)
	db.ORM.GetDB().AutoMigrate(
		&models.Task{},
	)

	// run the server
	fmt.Printf("Server is running on port %s", config.ServerConfigValues.Server.Port)
	addr := fmt.Sprint("127.0.0.1:" + config.ServerConfigValues.Server.Port)

	srv := &http.Server{
		// Addr: "0.0.0.0:8080",
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}

func setConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config.ServerConfigValues)
	if err != nil {
		log.Fatal(err)
	}
}
