package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func (app *Application) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	if td == nil {
		td = &TemplateData{}
	}

	td.CurrentYear = time.Now().Year()
	td.Flash = app.Session.PopString(r, "flash")
	td.IsAuthenticated = app.IsAuthenticated(r)
	td.CSRFToken = nosurf.Token(r)

	return td
}

func (app *Application) Render(w http.ResponseWriter, r *http.Request, name string, td *TemplateData) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("template %s does not exists", name))
		return
	}

	buf := new(bytes.Buffer)
	if err := ts.Execute(buf, app.AddDefaultData(td, r)); err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}
