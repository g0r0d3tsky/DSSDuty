package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/validator"
	"net/http"
)

func (app *app) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Course   int    `json:"course"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &domain.User{
		FullName:  input.Name,
		Email:     input.Email,
		Activated: false,
		Course:    input.Course,
	}
	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	v := validator.New()
	fmt.Printf(" %+v", user)

	err = app.UC.User.CreateUser(context.Background(), user)
	if err != nil {
		switch {

		case errors.Is(err, repository.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
