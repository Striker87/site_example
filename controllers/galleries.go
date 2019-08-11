package controllers

import (
	"site_example/models"
	"site_example/views"
)

type Galleries struct {
	New *views.View
	gs  models.GalleryService
}

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap.html", "galleries/new"),
		gs:  gs,
	}
}
