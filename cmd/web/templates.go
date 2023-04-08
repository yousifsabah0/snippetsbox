package main

import (
	"path/filepath"
	"text/template"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
)

type TemplateData struct {
	Snippets []*models.Snippet
	Snippet  *models.Snippet
}

type TemplateCache map[string]*template.Template

func NewTemplateCache(dir string) (TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
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
