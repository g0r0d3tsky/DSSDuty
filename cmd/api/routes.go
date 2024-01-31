package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *app) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	//router.ServeFiles("/static/*filepath", http.Dir("./ui/static"))

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/duties", app.createDutyHanler)
	router.HandlerFunc(http.MethodGet, "/v1/duties/:id", app.showDutyHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	return app.recoverPanic(router)
}
