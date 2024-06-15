package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// Create new ServeMux router.
	mux := chi.NewMux()
	// Use middlewares
	mux.Use(secureHeaders, app.LogRequest, app.recoverPanic)
	// Create a file server which serves static files.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register file server handler.
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	// Register Handler func.
	mux.Get("/", app.home)
	mux.Get("/snippet/view/{id}", app.snipperView)
	mux.Get("/snippet/create", app.snippetCreate)
	mux.Post("/snippet/create", app.snippetCreatePost)
	// Set notFound handler
	mux.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	}))

	return app.recoverPanic(app.LogRequest(secureHeaders(mux)))
}