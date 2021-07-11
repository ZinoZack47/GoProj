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
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
