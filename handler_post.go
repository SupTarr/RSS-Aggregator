package main

import (
	"fmt"
	"net/http"

	"github.com/SupTarr/RSS-Aggregator/internal/database"
)

func (apiCfg apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  20,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get posts: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
