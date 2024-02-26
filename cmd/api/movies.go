package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
	return
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		app.logger.Println(err.Error())
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie: %d\n", movieID)
	return
}
