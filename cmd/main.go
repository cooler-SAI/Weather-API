package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
	"weather-api/internal/api"
	"weather-api/internal/config"
	"weather-api/tools"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	cfg := config.Load()

	log.Info().
		Str("port", cfg.Port).
		Msg("Starting server...")

	log.Info().
		Str("url", "http://localhost:"+cfg.Port).
		Msg("Server URL")

	http.HandleFunc("/", api.WeatherHandler)

	go func() {
		if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error starting server")
		}
	}()

	tools.StopApp()
}
