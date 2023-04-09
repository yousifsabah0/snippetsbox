package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
	"github.com/yousifsabah0/snippetsbox/pkg/validators"
)

type TemplateData struct {
	CurrentYear int
	Flash       string
	Form        *validators.Form
	Snippets    []*models.Snippet
	Snippet     *models.Snippet
}

type TemplateCache map[string]*template.Template

func NormalizeDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"NormalizeDate": NormalizeDate,
}

func NewTemplateCache(dir string) (TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
