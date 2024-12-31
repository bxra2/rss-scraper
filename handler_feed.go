package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bxra2/rss-scraper/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameteres struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameteres{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	newUser, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user: %v", err))
		return
	}
	respondWithJSON(w, 201, newUser)
}

// func (apiCfg *apiConfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request, user database.User) {
// 	respondWithJSON(w, 200, databaseUsertoUser(user))
// }
