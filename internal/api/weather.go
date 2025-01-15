package api

import (
	"encoding/json"
	"net/http"
	"weather-api/internal/service"

	"github.com/rs/zerolog/log"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Moscow"
	}

	weatherData, err := service.GetWeatherData(city)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get weather data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weatherData); err != nil {
		log.Error().Err(err).Msg("Failed to encode response")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
