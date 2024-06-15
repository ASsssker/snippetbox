package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"snippetbox/internal/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// If url path is not root.
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprint(w, "%+v\n", snippet)
	}


	// // template files lists.
	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// }

	// // Parse of the template files.
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// // Render template.
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
}

// View snippets handler
func (app *application) snipperView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id from the url query parameter and
	// convert to integer.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// If an error is received during conversion or 
	// id < 1 return not found.
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

// Create new snippet handler
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// If method is not POST
	if r.Method != "POST" {
		// Set header and save 
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– KobayashiIssa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}