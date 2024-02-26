package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	movieID, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || movieID < 0 {
		return 0, fmt.Errorf("Error while reading id from params: %w", err)
	}
	return movieID, nil
}
