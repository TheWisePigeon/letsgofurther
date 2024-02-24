package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version = "1.0.0"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var appConfig config
	flag.IntVar(&appConfig.port, "port", 4000, "API server port")
	flag.StringVar(&appConfig.env, "env", "development", "Application environment (development|staging|production)")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: appConfig,
		logger: logger,
	}
	mux := http.NewServeMux()
	mux.Handle("GET /v1/healthcheck", nil)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("Starting %s server on port %s", app.config.env, server.Addr)
	err := server.ListenAndServe()
	logger.Fatal(err)
}
