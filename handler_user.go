package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	auth "github.com/Tusharpaul231/RSS-aggregator/internal/authentication"
	"github.com/Tusharpaul231/RSS-aggregator/internal/database"
	"github.com/google/uuid"
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
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     params.Name + "@example.com", // Placeholder email, should be replaced with actual logic
		Username:  params.Name, // Placeholder username, should be replaced with actual logic
		
	})
	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Failed to create user: ", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil{
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))

}

func databaseUserToUser(user any) interface{} {
	panic("unimplemented")
}