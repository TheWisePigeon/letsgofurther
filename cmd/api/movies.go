package main

import (
	"fmt"
	"greenlight/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
	return
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		app.logger.Println(err.Error())
		http.NotFound(w, r)
		return
	}
	movie := &data.Movie{
		ID:        movieID,
		CreatedAt: time.Now(),
		Title:     "Polar",
		Runtime:   69,
		Genres:    []string{"Action", "Violence", "Spy"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
