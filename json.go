package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	// respondWithError sends a JSON response with an error message and the specified HTTP status code.
	// It sets the Content-Type header to application/json and encodes the error message into JSON format.
	if code < 400 || code >= 600 {
		log.Printf("Invalid status code: %d. Must be between 400 and 599.", code)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}
	
	respondWithJSON(w, code, ErrorResponse{
		Error: message,
	})
}
	
	
// respondWithJSON sends a JSON response with the given status code and data.
// It sets the Content-Type header to application/json and encodes the data into JSON format.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshalling JSON:", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}