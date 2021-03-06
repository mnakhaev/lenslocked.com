package controllers

import (
	"lenslocked.com/views"
	"net/http"
)

type Galleries struct {
	NewView *views.View
}

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
