package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bxra2/rss-scraper/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameteres struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameteres{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	newUser, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user: %v", err))
		return
	}
	respondWithJSON(w, 200, newUser)
}
