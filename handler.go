package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Hello, World!",
	}

	log.Printf("success send response")
	json.NewEncoder(w).Encode(response)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]any{
		"data": users,
	}

	log.Printf("success send response")
	json.NewEncoder(w).Encode(response)
}
