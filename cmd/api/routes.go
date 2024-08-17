package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/ping", app.ping)

	return router
}

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
