package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%v", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func (app *Application) NotFoundError(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
