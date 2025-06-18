package main

import (
	"net/http"

	"github.com/Tusharpaul231/RSS-aggregator/internal/database"
	auth "github.com/Tusharpaul231/RSS-aggregator/internal/authentication"
)

type autheHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler autheHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the API key from the request header
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil{
			respondWithError(w, http.StatusUnauthorized, "Could not extract API key from request header")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Could not retrieve user by API key")
			return
	}

	handler(w, r, user)
	}
}