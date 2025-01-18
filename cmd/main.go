package main

import (
	"net/http"
	"os"
	"time"
	"weather-api/internal/api"
	"weather-api/internal/cache"
	"weather-api/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	cfg := config.Load()

	var cacheClient cache.Cache
	if cfg.CacheType == "redis" {
		cacheClient = cache.NewRedisCache(cfg.RedisAddr)
		log.Info().Msgf("Using Redis cache at %s", cfg.RedisAddr)
	} else {
		cacheClient = cache.NewMockCache()
		log.Info().Msg("Using Mock cache")
	}

	api.InitCache(cacheClient)

	log.Info().Msgf("Starting server... port=%s", cfg.Port)
	log.Info().Msgf("Server URL: http://localhost:%s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal().Err(err).
			Msg("Failed to start server")
	}
}
