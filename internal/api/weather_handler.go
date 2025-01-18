package api

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"weather-api/internal/cache"
	"weather-api/internal/service"

	"github.com/rs/zerolog/log"
)

var cacheClient cache.Cache

func InitCache(client cache.Cache) {
	cacheClient = client
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	expectedKey := os.Getenv("OPENWEATHER_API_KEY")

	if apiKey != expectedKey {
		log.Error().Str("received_key", apiKey).Msg("API key is missing or invalid")
		http.Error(w, `{"error":"API key is missing or invalid"}`, http.StatusInternalServerError)
		return
	}

	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Moscow"
	}

	cachedData, err := cacheClient.Get(city)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedData))
		return
	}

	weatherData, err := service.GetWeatherData(city)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch weather data")
		http.Error(w, `{"error":"Failed to fetch weather data"}`, http.StatusInternalServerError)
		return
	}

	weatherJSON, _ := json.Marshal(weatherData)

	if err := cacheClient.Set(city, string(weatherJSON), time.Hour); err != nil {
		log.Warn().Err(err).Msg("Failed to cache weather data")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(weatherJSON)
}
