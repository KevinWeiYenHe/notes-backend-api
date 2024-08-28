package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// a route to handle 405 METHOD NOT ALLOWED response
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/ping", app.pingHandler)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/notes", app.createNoteHandler)
	router.HandlerFunc(http.MethodGet, "/v1/notes/:id", app.showNoteHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/notes/:id", app.updateNoteHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/notes/:id", app.deleteNoteHandler)

	router.HandlerFunc(http.MethodGet, "/v1/notes", app.listNotesHandler)

	return router
}
