package main

import (
	"net/http"

	"github.com/ZinoZack47/GoProj.com/views"
	"github.com/gorilla/mux"
)

var (
	homeView     *views.View
	contactView  *views.View
	faqView      *views.View
	error404View *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, r); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, r); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqView.Template.Execute(w, r); err != nil {
		panic(err)
	}
}

func error404notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := error404View.Template.Execute(w, r); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	faqView = views.NewView("views/faq.gohtml")
	error404View = views.NewView("views/error404.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(error404notfound)
	http.ListenAndServe(":3000", r)
}
