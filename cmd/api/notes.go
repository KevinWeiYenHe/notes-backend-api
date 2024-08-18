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
	w.Write([]byte("Creating a Note"))
}

func (app *application) showNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get param id
	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	//
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
		fmt.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
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
