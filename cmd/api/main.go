package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

// hardcoded app version number
const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	// Possible command line flags to be called
	flag.IntVar(&cfg.port, "addr", 4000, "Server address")
	flag.StringVar(&cfg.env, "env", "development", "Server type")

	flag.Parse()

	app := &application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Launching %s server on port %d...", app.config.env, app.config.port)
	err := srv.ListenAndServe()
	fmt.Println(err)

}
