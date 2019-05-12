package controllers

import (
	"fmt"
	"log"
	"net/http"
	"tensor_web_application/views"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap.html", "users/new"),
	}
}

// This is used to render the form where a user can create a new users account
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// using when user try to create a user account from HTML form
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	err := parseForm(r, &form)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, form)
}

