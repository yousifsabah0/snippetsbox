package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	handler := nosurf.New(next)
	handler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return handler
}

func (app *Application) IsAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r, "user_id")
}

func (app *Application) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.IsAuthenticated(r) {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (app *Application) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.InfoLogger.Printf("%s - %s %s %s", r.RemoteAddr, r.Method, r.Proto, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RevcoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.ServerError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
