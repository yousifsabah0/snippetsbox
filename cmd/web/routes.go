package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *Application) Routes() http.Handler {

	middlewares := alice.New(app.RevcoverPanic, app.LogRequest, SecureHeaders)

	// Initialize a new servemux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippets", app.ShowSnippet)
	mux.HandleFunc("/snippets/new", app.CreateSnippet)

	// Serve static files, e.g (stylesheets, javascript, and images)
	fileserver := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	return middlewares.Then(mux)
}
