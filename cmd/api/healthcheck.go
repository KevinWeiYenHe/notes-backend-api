package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, h *http.Request) {
	data := fmt.Sprintf(`{"status" : "available", "environment" : "%s", "version" : "%s"}`, app.config.env, version)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}
