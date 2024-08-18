package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

// writing JSON out
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Convert data to json byte data
	// js, err := json.Marshal(data)
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// add new line to make it prettier for terminal users
	js = append(js, '\n')

	// Add headers that we want to add to the response
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
