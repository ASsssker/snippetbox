package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// addr is a pointer to command line option with the nme "addr"
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parsing command line options
	flag.Parse()
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
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}