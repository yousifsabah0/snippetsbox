package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const (
	BaseTemplateFile string = "./web/html/base.layout.html"
	HomeTemplateFile string = "./web/html/pages/home.page.html"
	ShowTemplateFile string = "./web/html/pages/show.page.html"

	FooterPartialFile string = "./web/html/partials/footer.partial.html"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := app.Snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippets?id=%d", id), http.StatusPermanentRedirect)
}
