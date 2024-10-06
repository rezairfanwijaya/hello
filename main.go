package main

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/hello.git/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/users", handler.HandleUser)

	log.Println("starting the server on :8181")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatalf("failed run server on port 8181, err: %v", err)
	}
}
