package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *Application) Routes() http.Handler {

	middlewares := alice.New(app.RevcoverPanic, app.LogRequest, SecureHeaders)

	// Initialize a new servemux
	mux := pat.New()

	// Register routes
	mux.Get("/", http.HandlerFunc(app.Home))

	mux.Get("/snippets/new", http.HandlerFunc(app.CreateSnippetForm))
	mux.Post("/snippets/new", http.HandlerFunc(app.CreateSnippet))

	mux.Get("/snippets/:id", http.HandlerFunc(app.ShowSnippet))

	// Serve static files, e.g (stylesheets, javascript, and images)
	fileserver := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	return middlewares.Then(mux)
}
