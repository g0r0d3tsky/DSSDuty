package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *app) createDutyHanler(w http.ResponseWriter, r *http.Request) {
	//placeholder
	var input struct {
		Date   time.Time `json:"date"`
		Amount int       `json:"amount"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	// Dump the contents of the input struct in a HTTP response.
	fmt.Fprintf(w, "%+v\n", input)
}
func (app app) showDutyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		log.Fatal("invalid duty id")
		return
	}
	duty, err := app.UC.Duty.GetDutyByID(context.Background(), id)
	if err != nil {
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"duty": duty}, nil)
	if err != nil {
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}
