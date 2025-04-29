// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"urlshortenergo/handlers"
	"urlshortenergo/storage"

	"github.com/gorilla/mux"
)

func main() {
	rdb := storage.NewRedisClient("localhost:6379")

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	s := handlers.NewServer(rdb, baseURL)
	r := mux.NewRouter()
	r.HandleFunc("/shorten", s.ShortenHandler).Methods("POST")
	r.HandleFunc("/{code}", s.RedirectHandler).Methods("GET")

	fmt.Println("Server running on", baseURL)
	log.Fatal(http.ListenAndServe(":8080", r))
}
