package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	"D:\My-Projects\RSS-aggregator\internal\database\database.go"
)



func (apiCfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	// This handler is used to create a new user.
	// It should parse the request body, validate the input, and create a new user in the database.
	// For now, it responds with a 201 Created status and a placeholder message.
	type parameters struct {
		Name  string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:   uuid.New(),
		created_at: time.Now().UTC(),
		updated_at: time.Now().UTC(),
		Name: params.Name,
	})
	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}