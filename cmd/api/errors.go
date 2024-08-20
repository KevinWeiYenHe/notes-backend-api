package main

import (
	"fmt"
	"net/http"
)

// Generic helper function for logging an error message
func (app *application) logError(r *http.Request, err error) {
	fmt.Println(err)
}

// handles errors and respond back to the user as a JSON message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// 400 BAD REQUEST
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// 404 NOT FOUND
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resoruce could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)

}

// 405 METHOD NOT ALLOWED
// handles issues where a client has made a request where the method is not supported for that resource
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// 422 UNPROCESSABLE ENTITY
// Note that the errors parameter here has the type map[string]string, which is exactly
// the same as the errors map contained in our Validator type.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

// 500 INTERNAL SERVER ERROR
// handles errors that occur at the server level
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "The server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}
