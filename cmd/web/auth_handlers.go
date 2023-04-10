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

	form := validators.New(r.PostForm)
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
	app.Render(w, r, "login.page.html", &TemplateData{
		Form: validators.New(nil),
	})
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := validators.New(r.PostForm)
	form.Required("email", "password")
	form.MatchPattern("email", validators.EmailRX)
	form.MinLength("password", 8)

	if !form.Valid() {
		app.Render(w, r, "login.page.html", &TemplateData{
			Form: form,
		})
	}

	id, err := app.Users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or password incorrect!")
			app.Render(w, r, "login.page.html", &TemplateData{
				Form: form,
			})
		} else {
			app.ServerError(w, err)
		}
		return
	}

	app.Session.Put(r, "user_id", id)
	http.Redirect(w, r, "/snippets/new", http.StatusSeeOther)
}

func (app *Application) Logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "user_id")
	app.Session.Put(r, "flash", "You've been logged out!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
