package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to the Site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(w, "Contact us <a href=\"mailto:martinkalvin6@gmail.com\">our email</a> ")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `<h1>ERROR 404: The page you are trying to reach does not exist.</h1>
		<p>if you keep getting this page please contact us at out email.</p>`)
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
