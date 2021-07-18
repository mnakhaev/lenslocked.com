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
		ContactView: views.NewView("bootstrap", "views/static/contact.gohtml"),
		HomeView:    views.NewView("bootstrap", "views/static/home.gohtml"),
		FaqView:     views.NewView("bootstrap", "views/static/faq.gohtml"),
	}
}
