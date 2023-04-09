package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *Application) Routes() http.Handler {

	middlewares := alice.New(app.RevcoverPanic, app.LogRequest, SecureHeaders)
	dynamicMiddlewares := alice.New(app.Session.Enable)

	// Initialize a new servemux
	mux := pat.New()

	// Register routes
	mux.Get("/", dynamicMiddlewares.ThenFunc(app.Home))

	mux.Get("/snippets/new", dynamicMiddlewares.ThenFunc(app.CreateSnippetForm))
	mux.Post("/snippets/new", dynamicMiddlewares.ThenFunc(app.CreateSnippet))

	mux.Get("/snippets/:id", dynamicMiddlewares.ThenFunc(app.ShowSnippet))

	mux.Get("/users/signup", dynamicMiddlewares.ThenFunc(app.SignupForm))
	mux.Get("/users/login", dynamicMiddlewares.ThenFunc(app.LoginForm))

	mux.Post("/users/signup", dynamicMiddlewares.ThenFunc(app.Signup))
	mux.Post("/users/login", dynamicMiddlewares.ThenFunc(app.Login))

	// Serve static files, e.g (stylesheets, javascript, and images)
	fileserver := http.FileServer(http.Dir("./web/static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileserver))

	return middlewares.Then(mux)
}
