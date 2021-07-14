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
	w.Header().Set("Content-Type", "text/html;charset=utf8")
	must(homeView.Render(w, r))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf8")
	must(contactView.Render(w, r))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf8")
	must(faqView.Render(w, r))
}

func error404notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf8")
	w.WriteHeader(http.StatusNotFound)
	must(error404View.Render(w, r))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	error404View = views.NewView("bootstrap", "views/error404.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(error404notfound)

	r.PathPrefix("/css/").
		Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./views/layouts/css"))))
	r.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./views/layouts/assets"))))

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
