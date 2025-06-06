package main

import (
	"net/http"
)

// handlerReadiness checks if the service is ready to handle requests.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// Respond with a 200 OK status to indicate readiness
	respondWithJSON(w, http.StatusOK, struct{}{})
}