package main

import (
	"encoding/json"
	"log"
	"net/http"
)


// hardcoded values 
const (
	apiKey  = "DUMMY_API_KEY_1234567890"
	userID  = "admin_demo_001"
	password = "DemoPassword123!"
	databaseURL = "postgres://demo_user:demo_pass@localhost:5432/demo"
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
	http.HandleFunc("/health", healthHandler)

	log.Println("Demo API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}