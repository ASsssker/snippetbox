package main

import (
	"log"
	"net/http"
)

func main() {
	// Create new ServeMux router.
	mux := http.NewServeMux()
	// Create a file server which serves static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register file server handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Register Handler func.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snipperView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// Use ListenAndServe func to start a new web server.
	// We pass in tcp socket address and ServeMux router.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}