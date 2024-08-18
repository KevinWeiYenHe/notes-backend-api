package main

import (
	"fmt"
	"net/http"
)

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"ping": "pong",
		"foo":  "bar",
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
