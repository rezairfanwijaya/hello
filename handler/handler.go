package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rezairfanwijaya/hello.git/data"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Hello, World!",
	}

	log.Printf("success send response")
	json.NewEncoder(w).Encode(response)
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]any{
		"data": data.Users,
	}

	log.Printf("success send response")
	json.NewEncoder(w).Encode(response)
}
