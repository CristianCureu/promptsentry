package main

import (
	"log"
	"net/http"

	"github.com/cristiancureu/prompt-sentry/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("API running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
