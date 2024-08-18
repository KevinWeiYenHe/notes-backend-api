package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/KevuTheDev/notes-backend-api/internal/data"
	"github.com/julienschmidt/httprouter"
)

func (app *application) createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`             // title of note
		Content string   `json:"content,omitempty"` // content of note
		Tags    []string `json:"tags,omitempty"`    // tags of note
	}

	// Decode the given body from the response, and store the value in ^input
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get param id
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResource(w, r)
		return
	}

	// The output data struct
	note := data.Note{
		ID:           id,
		CreatedAt:    time.Now(),
		LastUpdateAt: time.Now(),
		Title:        "Hello World",
		Content:      "This is my first note",
		Tags:         []string{"Message", "Hello", "World"},
		Version:      1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"note": note}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
