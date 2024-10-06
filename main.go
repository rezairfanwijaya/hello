package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", handleUser)

	log.Println("starting the server on :8181")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatalf("failed run server on port 8181, err: %v", err)
	}
}
