package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// If url path is not root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello, from Snippetbox"))
}

// View snippets handler
func snipperView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Create new snippet handler
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// If method is not POST
	if r.Method != "POST" {
		// Set header and save 
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Create new ServeMux router
	// and register Handler func.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snipperView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// Use ListenAndServe func to start a new web server.
	// We pass in tcp socket address and ServeMux router.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}