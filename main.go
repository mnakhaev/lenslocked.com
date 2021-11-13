package main

import (
	"fmt"
	"lenslocked.com/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>Current page is not yet implemented...</h1>")
}

func main() {
	galleriesC := controllers.NewGalleries()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	// Handle accepts object (suitable for static pages handling), HandleFunc accepts function.
	// For static views below it's only required to set content-type and render view.
	// That's why it's not necessary to use HandleFunc there.
	r.Handle("/", staticC.HomeView).Methods("GET")
	r.Handle("/contact", staticC.ContactView).Methods("GET")
	r.Handle("/faq", staticC.FaqView).Methods("GET")

	r.HandleFunc("/galleries/new", galleriesC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	_ = http.ListenAndServe(":8080", r)
}
