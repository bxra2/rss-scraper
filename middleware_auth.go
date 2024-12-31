package main

import (
	"fmt"
	"net/http"

	"github.com/bxra2/rss-scraper/internal/auth"
	"github.com/bxra2/rss-scraper/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth err %v: ", err))
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt get user %v: ", err))
			return
		}
		handler(w, r, user)
	}
}
