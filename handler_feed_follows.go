package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bxra2/rss-scraper/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameteres struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameteres{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	FeedFollows, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create Feed Follows: %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedFollowtoFeedFollow(FeedFollows))
}

func (apiCfg *apiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed Follows: %v", err))
	}
	respondWithJSON(w, 200, dbFeedFollowstoFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parse feed_follow id : %v", err))

	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(),
		database.DeleteFeedFollowsParams{
			ID:     feedFollowID,
			UserID: user.ID,
		})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed Follows: %v", err))
	}

	respondWithJSON(w, 200, struct {
		Success bool `json:"success"`
	}{
		Success: true,
	})
}
