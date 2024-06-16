package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// Create new ServeMux router.
	mux := chi.NewMux()
	// Create a file server which serves static files.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register Handler func.
	mux.Group(func(r chi.Router) {
		r.Use(secureHeaders, app.LogRequest, app.recoverPanic)
		r.Handle("/static/*", http.StripPrefix("/static", fileServer))

		r.Group(func(r chi.Router) {
			r.Use(app.sessionManager.LoadAndSave)
			r.Get("/", app.home)
			r.Get("/snippet/view/{id}", app.snipperView)
			r.Get("/snippet/create", app.snippetCreate)
			r.Post("/snippet/create", app.snippetCreatePost)

			r.Get("/user/signup", app.userSignup)
			r.Post("/user/signup", app.userSignupPost)

			r.Get("/user/login", app.userLogin)
			r.Post("/user/login", app.userLoginPost)

			r.Post("/user/logout", app.userLogoutPost)
		})
	})

	// Set notFound handler
	mux.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	}))

	return mux
}