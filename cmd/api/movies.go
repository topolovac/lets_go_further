package main

import (
	"fmt"
	"net/http"
	"time"

	"matejtopolovac.lgf/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := &data.Movie{
		ID:          id,
		CreatedTime: time.Now(),
		Title:       "My movie",
		Runtime:     1,
		Genres:      []string{"Drama", "Romance", "War"},
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
