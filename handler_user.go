package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alyhenr/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	var params paramaters
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Bad request")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create new user, err: %v", err))
		return
	}

	respondWithJson(w, 201, convertToResponseFormat(user))
}
