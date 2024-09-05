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

	router.HandlerFunc(http.MethodGet, "/v1/notes", app.listNotesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/notes", app.createNoteHandler)
	router.HandlerFunc(http.MethodGet, "/v1/notes/:id", app.showNoteHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/notes/:id", app.updateNoteHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/notes/:id", app.deleteNoteHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// temporary v2 API with user authenticated
	// temporary named v2 to keep legacy v1, as we incrementally update the frontend
	router.HandlerFunc(http.MethodGet, "/v2/ping", app.pingHandler)
	router.HandlerFunc(http.MethodGet, "/v2/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v2/notes", app.requireActivatedUser(app.listNotesByUserHandler))
	router.HandlerFunc(http.MethodPost, "/v2/notes", app.requireActivatedUser(app.createNoteByUserHandler))
	router.HandlerFunc(http.MethodGet, "/v2/notes/:id", app.requireActivatedUser(app.showNoteByUserHandler))
	router.HandlerFunc(http.MethodPatch, "/v2/notes/:id", app.requireActivatedUser(app.updateNoteByUserHandler))
	router.HandlerFunc(http.MethodDelete, "/v2/notes/:id", app.requireActivatedUser(app.deleteNoteByUserHandler))

	router.HandlerFunc(http.MethodPost, "/v2/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v2/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v2/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.authenticate(router))
}
