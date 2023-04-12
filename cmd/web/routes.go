package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *Application) Routes() http.Handler {

	middlewares := alice.New(app.RevcoverPanic, app.LogRequest, SecureHeaders)
	dynamicMiddlewares := alice.New(app.Session.Enable, NoSurf, app.Authenticate)

	// Initialize a new servemux
	mux := pat.New()

	// Register routes
	mux.Get("/ping", http.HandlerFunc(Ping))

	mux.Get("/", dynamicMiddlewares.Append(app.RequireAuthentication).ThenFunc(app.Home))

	mux.Get("/snippets/new", dynamicMiddlewares.Append(app.RequireAuthentication).ThenFunc(app.CreateSnippetForm))
	mux.Post("/snippets/new", dynamicMiddlewares.Append(app.RequireAuthentication).ThenFunc(app.CreateSnippet))

	mux.Get("/snippets/:id", dynamicMiddlewares.Append(app.RequireAuthentication).ThenFunc(app.ShowSnippet))

	mux.Get("/users/signup", dynamicMiddlewares.ThenFunc(app.SignupForm))
	mux.Get("/users/login", dynamicMiddlewares.ThenFunc(app.LoginForm))

	mux.Post("/users/signup", dynamicMiddlewares.ThenFunc(app.Signup))
	mux.Post("/users/login", dynamicMiddlewares.ThenFunc(app.Login))

	mux.Post("/users/logout", dynamicMiddlewares.Append(app.RequireAuthentication).ThenFunc(app.Logout))

	// Serve static files, e.g (stylesheets, javascript, and images)
	fileserver := http.FileServer(http.Dir("./web/static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileserver))

	return middlewares.Then(mux)
}
