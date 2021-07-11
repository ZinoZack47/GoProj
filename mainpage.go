package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to the Site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Contact us <a href=\"mailto:martinkalvin6@gmail.com\">our email</a>")
}

/*fmt.Fprint(w, `<h1>ERROR 404: The page you are trying to reach does not exist.</h1>
<p>if you keep getting this page please contact us at out email.</p>`)*/

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
