package main

import (
	"log"
	"net/http"
	"os"

	"Weather-API/internal/api"
	"Weather-API/internal/config"
)

func main() {
	cnf := config.Load()

	http.HandleFunc("/", api.WeatherHandler)

	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Error to start Server: %v", err)
	}

}
