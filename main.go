package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "Hello, World!",
		}

		json.NewEncoder(w).Encode(response)
	})

	log.Println("starting the server on :8181")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatalf("failed run server on port 8181, err: %v", err)
	}
}
