package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"site_example/controllers"
	"site_example/models"
	"site_example/views"
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

	services, err := models.NewServices(psqlInfo)
	must(err)

	defer services.Close()
	//services.DestructiveReset()
	services.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	//r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
