package main

import "net/http"

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {

	data := `{"ping" : "pong"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}
