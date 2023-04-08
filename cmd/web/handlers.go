package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const (
	BaseTemplateFile string = "./web/html/base.layout.html"
	HomeTemplateFile string = "./web/html/home.page.html"

	FooterPartialFile string = "./web/html/partials/footer.partial.html"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		HomeTemplateFile,
		BaseTemplateFile,
		FooterPartialFile,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.NotFoundError(w)
		return
	}

	fmt.Fprintf(w, "Dispaly a specific snippet with id %d!", id)
}

func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("New snippet is created!"))
}
