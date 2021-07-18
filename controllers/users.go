package controllers

import (
	"fmt"
	"lenslocked.com/views"
	"net/http"
)

type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := &SignUpForm{}
	if err := parseForm(r, form); err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintln(w, "Email is", form.Email)
	_, _ = fmt.Fprintln(w, "Password is", form.Password)
}
