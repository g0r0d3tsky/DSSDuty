package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"net/http"
	"strings"
)

func (app *app) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *app) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			app.invalidCredentialsResponse(w, r)
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}
		token := headerParts[1]

		if len(token) == 0 {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		userID, err := app.UC.Auth.ParseToken(token)
		if err != nil {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}
		user, err := app.UC.User.GetUserByID(context.Background(), userID)
		if err != nil {
			app.badRequestResponse(w, r, err)
		}
		ctx := context.WithValue(r.Context(), "user", user)
		r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (app *app) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*domain.User)
		if !ok {
			app.badRequestResponse(w, r, errors.New("missing user value in request context"))
		}

		if !user.Activated {
			app.inactiveAccountResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *app) requireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*domain.User)
		if !ok {
			app.badRequestResponse(w, r, errors.New("missing user value in request context"))
		}

		if user.Role != domain.ADMIN {
			app.methodNotAllowedResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
