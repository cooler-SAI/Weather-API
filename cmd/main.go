package main

import (
	"net/http"
	"os"
	"time"
	"weather-api/internal/api"
	"weather-api/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	// log configs
	cfg := config.Load()

	// log server start
	log.Info().
		Str("port", cfg.Port).
		Msg("Starting server...")

	// added handler
	http.HandleFunc("/", api.WeatherHandler)

	// start server
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error starting server")
	}
}
