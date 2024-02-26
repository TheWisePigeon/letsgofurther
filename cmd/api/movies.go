package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
	return
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	movieID, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || movieID < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie: %d\n", movieID)
	return
}
