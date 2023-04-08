package main

import (
	"fmt"
	"net/http"
)

func (app *Application) Render(w http.ResponseWriter, r *http.Request, name string, td *TemplateData) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("template %s does not exists", name))
		return
	}

	if err := ts.Execute(w, td); err != nil {
		app.ServerError(w, err)
	}
}
