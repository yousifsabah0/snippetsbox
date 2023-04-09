package main

import (
	"errors"
	"net/http"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
	"github.com/yousifsabah0/snippetsbox/pkg/validators"
)

func (app *Application) SignupForm(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "signup.page.html", &TemplateData{
		Form: validators.New(nil),
	})
}

func (app *Application) Signup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := validators.New(r.Form)
	form.Required("name", "email", "password")
	form.MatchPattern("email", validators.EmailRX)
	form.MinLength("form", 8)

	if !form.Valid() {
		app.Render(w, r, "signup.page.html", &TemplateData{
			Form: form,
		})
	}

	if err := app.Users.Insert(form.Get("name"), form.Get("email"), form.Get("password")); err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Email address is already in use!")
			app.Render(w, r, "signup.page.html", &TemplateData{
				Form: form,
			})
		} else {
			app.ServerError(w, err)
		}
		return
	}

	app.Session.Put(r, "flash", "Account created!")

	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}

func (app *Application) LoginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TOBE"))
	// app.Render(w, r, "login.page.html", &TemplateData{
	// 	Form: validators.New(nil),
	// })
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {}
