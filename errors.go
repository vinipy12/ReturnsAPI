package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(ErrorResponse{Error: msg}); err != nil {
		log.Printf("Failed to encode error: %v", err)
	}
}
