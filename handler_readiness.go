package main

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if(statusCode> 499){
		log.Println("Responding with 5XX Error: ", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, statusCode, errorResponse{Error: message})
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w,200,struct{}{})
}