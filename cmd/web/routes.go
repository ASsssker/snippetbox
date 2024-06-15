package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Create new ServeMux router.
	mux := http.NewServeMux()
	// Create a file server which serves static files.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register file server handler.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Register Handler func.
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snipperView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}