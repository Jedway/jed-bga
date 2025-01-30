package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

// Response struct for JSON output
type Response struct {
	Email         string `json:"email"`
	CurrentTime   string `json:"current_datetime"`
	GitHubRepoURL string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Email:         "your-email@example.com",
		CurrentTime:   time.Now().UTC().Format(time.RFC3339),
		GitHubRepoURL: "https://github.com/yourusername/go-public-api",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// CORS Middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal(err)
	}
}
