package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tensor_web_application/controllers"
	"tensor_web_application/models"
	"tensor_web_application/views"
)

var tpl *views.View

func must(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "test_site"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	must(err)
	defer us.Close()

	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/login", usersC.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
