package main

import (
	"net/http"
	"time"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/tharindu1998/go-crud/internal/database"
)
func(apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name  string `json:"name"`
	}
	decoder:= json.NewDecoder(r.Body)
	params:= parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid request payload")
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID :uuid.New(),
		Username: 	params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		
	})
	if err != nil {
	// Log the detailed DB error
	http.Error(w, "Couldn't create user", http.StatusInternalServerError)
	// Use log.Printf or your logger of choice
	log.Printf("CreateUser error: %v\n", err)
	return
	}


	respondWithJSON(w, http.StatusOK, user)
	
}