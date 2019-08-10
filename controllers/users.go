package controllers

import (
	"fmt"
	"log"
	"net/http"
	"site_example/models"
	"site_example/rand"
	"site_example/views"
)

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        models.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap.html", "users/new"),
		LoginView: views.NewView("bootstrap.html", "users/login"),
		us:        us,
	}
}

// This is used to render the form where a user can create a new users account
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {

	//a := Alert{
	//	Level:   "success",
	//	Message: "Successfully rendered a dynamic alert!",
	//}
	//d := Data{
	//	Alert: a,
	//	Yield: "Hello",
	//}
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

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, form)

	err = u.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address")
		case models.ErrPasswordIncorrect:
			fmt.Fprintln(w, "Invalid password provided")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, user)
}

// CookieTest is used to display cookie set on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, user)
}

func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	return nil
}
