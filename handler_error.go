package main

import (
	"net/http"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	// This handler is used to simulate an error response.
	// It responds with a 500 Internal Server Error status code and a JSON error message.
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}