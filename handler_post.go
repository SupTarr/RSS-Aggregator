package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SupTarr/RSS-Aggregator/internal/database"
)

func (apiCfg apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitInt := 20
	limit := r.URL.Query()["limit"]
	pageInt := 1
	page := r.URL.Query()["page"]

	if len(limit) > 0 {
		l, err := strconv.Atoi(limit[0])
		if err != nil {
			log.Printf("Cannot convert %v to an int", limit)
		}

		limitInt = l
	}

	if len(page) > 0 {
		p, err := strconv.Atoi(page[0])
		if err != nil {
			log.Printf("Cannot convert %v to an int", page)
		}

		pageInt = p
	}

	offset := (pageInt - 1) * limitInt

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limitInt),
		Offset: int32(offset),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get posts: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
