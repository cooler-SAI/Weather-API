package main

import (
	"log"
	"net/http"
	"weather-api/internal/api"
	"weather-api/internal/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/", api.WeatherHandler)

	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
