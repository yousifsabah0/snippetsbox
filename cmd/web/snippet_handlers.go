package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/yousifsabah0/snippetsbox/pkg/validators"
)

const (
	BaseTemplateFile string = "./web/html/base.layout.html"
	HomeTemplateFile string = "./web/html/pages/home.page.html"
	ShowTemplateFile string = "./web/html/pages/show.page.html"

	FooterPartialFile string = "./web/html/partials/footer.partial.html"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.Snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.Render(w, r, "home.page.html", &TemplateData{
		Snippets: snippets,
	})
}

func (app *Application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		app.NotFoundError(w)
		return
	}

	snippet, err := app.Snippets.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.NotFoundError(w)
		} else {
			app.ServerError(w, err)
		}
	}

	app.Render(w, r, "show.page.html", &TemplateData{
		Snippet: snippet,
	})
}

func (app *Application) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "create.page.html", &TemplateData{
		Form: validators.New(nil),
	})
}

func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.ClientError(w, http.StatusBadRequest)
	}

	form := validators.New(r.Form)
	form.Required("title", "content", "expires")
	form.Length("title", 100)
	form.PermittedValues("expires", "365", "6", "1")

	if !form.Valid() {
		app.Render(w, r, "create.page.html", &TemplateData{
			Form: form,
		})
		return
	}

	id, err := app.Snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippets/%d", id), http.StatusSeeOther)
}
