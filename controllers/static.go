package controllers

import "lenslocked.com/views"

// Static represents views for static pages
type Static struct {
	ContactView *views.View
	HomeView    *views.View
	FaqView     *views.View
}

func NewStatic() *Static {
	return &Static{
		ContactView: views.NewView("bootstrap", "static/contact"),
		HomeView:    views.NewView("bootstrap", "static/home"),
		FaqView:     views.NewView("bootstrap", "static/faq"),
	}
}
