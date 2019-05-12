package controllers

import "tensor_web_application/views"

type Static struct {
	Home *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap.html", "static/home"),
		Contact: views.NewView("bootstrap.html", "static/contact"),
	}
}