package main

import (
	"context"
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (app *app) createDutyHanler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Date   time.Time `json:"date"`
		Amount int       `json:"amount"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	//TODO: delete
	fmt.Printf("value: %+v \n", input)

	duty := &domain.Duty{
		Date:   input.Date,
		Amount: input.Amount,
	}

	err = app.UC.Duty.CreateDuty(context.Background(), duty)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"duty": duty}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app app) pickDutyHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		dutyID uuid.UUID `json:"duty_id"`
		time   time.Time `json:"time"`
		userID uuid.UUID `json:"user_id"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

}

func (app app) showAvailableDutyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		log.Fatal("invalid user id")
		return
	}
	duties, err := app.UC.Duty.GetAvailableDuty(context.Background(), id)
	if err != nil {
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"duties": duties}, nil)
	if err != nil {
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
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
