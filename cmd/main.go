package main

import (
	"net/http"
	"os"
	"time"
	"weather-api/internal/api"
	"weather-api/internal/cache"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	err := godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg("No .env file found, using system environment variables")
	}

	cacheType := os.Getenv("CACHE_TYPE")
	var cacheClient cache.Cache
	if cacheType == "redis" {
		redisAddr := os.Getenv("REDIS_ADDR")
		if redisAddr == "" {
			redisAddr = "localhost:6379"
		}
		cacheClient = cache.NewRedisCache(redisAddr)
		log.Info().Msg("Using Redis cache")
	} else {
		cacheClient = cache.NewMockCache()
		log.Info().Msg("Using Mock cache")
	}

	api.InitCache(cacheClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", api.WeatherHandler)

	log.Info().Msgf("Starting server... port=%s", port)
	log.Info().Msgf("Server URL: http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
