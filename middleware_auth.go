package main

import (
	"fmt"
	"net/http"

	"github.com/SupTarr/RSS-Aggregator/internal/database"
	"github.com/SupTarr/RSS-Aggregator/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not get user: %s", err))
			return
		}

		handler(w, r, user)
	}
}
