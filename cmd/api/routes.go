package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/ping", ping)

	return router
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
