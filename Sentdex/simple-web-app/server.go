package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Build a handler
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First response")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r)
}
