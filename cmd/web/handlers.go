package main

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (app *Application) About(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "about.page.html", nil)
}
