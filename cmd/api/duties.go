package main

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
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
	//layout := "2006-01-02 15:04:05"
	//date, err := time.Parse(layout, dateStr)
	if err != nil {
		// Обработка ошибки парсинга даты
	} else {
		// Используйте переменную date для работы с датой
	}
	duty := &domain.Duty{
		//Date:   ,
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
