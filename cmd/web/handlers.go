package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	// Extract the value of the id from the url query parameter and
	// convert to integer.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// If an error is received during conversion or 
	// id < 1 return not found
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
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