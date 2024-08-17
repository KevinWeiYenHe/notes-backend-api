package main

import (
	"flag"
	"fmt"
	"net/http"
)

type application struct {
}

type config struct {
	Addr string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.Addr, "addr", ":4000", "Server address")

	app := application{}

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	fmt.Println(err)

}
