package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Email           string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GithubURL       string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// setting CORS headers for CORS handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// preparing response
	response := Response{
		Email:           "jedotoekong@gmail.com",                                  // hng registered email
		CurrentDatetime: time.Now().UTC().Add(1 * time.Hour).Format(time.RFC3339), //  current datetime in ISO 8601 format (yymmdd h-m-s-ms)
		GithubURL:       "https://github.com/Jedway/jed-bga.git",                  // my github profile
	}

	// encode the response in JSON format
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
