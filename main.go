package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from Snippetbox"))
}

func main() {
	// Create new ServeMux router
	// and register Handler func.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// Use ListenAndServe func to start a new web server.
	// We pass in tcp socket address and ServeMux router.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}