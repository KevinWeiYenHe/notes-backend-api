package main

import (
	"net/http"
)

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"ping": "pong",
		"foo":  "bar",
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
