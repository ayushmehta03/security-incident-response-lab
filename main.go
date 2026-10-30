package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(
		HealthResponse{Status: "ok"},
	)
}

func main() {

	//fixing the hard coded part 
	
	apiKey := os.Getenv("API_KEY")
	databaseURL := os.Getenv("DATABASE_URL")

	if apiKey == "" || databaseURL == "" {
		log.Fatal("Required application configuration is missing")
	}

	_ = apiKey
	_ = databaseURL

	http.HandleFunc("/health", healthHandler)

	log.Println("Demo API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}