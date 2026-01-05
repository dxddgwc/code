package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(helloResponse{Message: "Hello, World!"}); err != nil {
			log.Printf("ERROR: could not encode response: %v", err)
		}
	})

	log.Println("INFO: server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("FATAL: could not start server: %v", err)
	}
}
