package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tensor_web_application/controllers"
	"tensor_web_application/views"
)

var tpl *views.View

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
