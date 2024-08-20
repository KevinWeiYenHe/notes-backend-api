package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KevuTheDev/notes-backend-api/internal/data"
	"github.com/KevuTheDev/notes-backend-api/internal/validator"
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

	// copy the values from the input struct to a new Note struct
	note := &data.Note{
		Title:   input.Title,
		Content: input.Content,
		Tags:    input.Tags,
	}

	// Initialize a new Validator
	v := validator.New()
	// Perform validation check on data sent from client
	if data.ValidateNote(v, note); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// validation check passed, performing insert
	err = app.models.Notes.Insert(note)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	// setup a location header of where the resource will be located at
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/notes/%d", note.ID))

	// send a response back to the client with the new note
	err = app.writeJSON(w, http.StatusCreated, envelope{"note": note}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// obtain the id portion of the router
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) showNoteHandler(w http.ResponseWriter, r *http.Request) {
	// get id param from the URI
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// get note based on id (extracted from URI)
	note, err := app.models.Notes.Get(id)
	if err != nil {
		switch {
		// no record found of specified id
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		// any errors that occur in the process of obtaining record
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// send a response of the obtained note
	err = app.writeJSON(w, http.StatusOK, envelope{"note": note}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	// get id param from the URI
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// get note specified by id
	// get note to see if the note exists in the database
	// if exists, proceed to use this data and then update it provided by client
	note, err := app.models.Notes.Get(id)
	if err != nil {
		switch {
		// no record found of specified id
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		// any errors that occur in the process of obtaining record
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	// Decode the given body from the response, and store the value in ^input
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// reuse the struct
	note.Title = input.Title
	note.Content = input.Content
	note.Tags = input.Tags

	// Initialize a new Validator
	v := validator.New()
	// Perform validation check on data sent from client
	if data.ValidateNote(v, note); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Perform an update on the given data
	err = app.models.Notes.Update(note)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// when successful, create a request to user with the new note data
	err = app.writeJSON(w, http.StatusOK, envelope{"note": note}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// get id param from the URI
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Perform a delete on record based on id
	err = app.models.Notes.Delete(id)
	if err != nil {
		switch {
		// no record found of specified id
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		// any errors that occur in the process of obtaining record
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// if delete record was possible, send message of successful deletion
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully delete"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
