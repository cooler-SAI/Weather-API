package api

import (
	"encoding/json"
	"net/http"
	"time"
	"weather-api/internal/cache"
	"weather-api/internal/service"

	"github.com/rs/zerolog/log"
)

var redisClient = cache.NewRedisClient("localhost:6379")

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Moscow"
	}

	cachedData, err := redisClient.Get(r.Context(), city).Result()
	if err == nil {

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(cachedData))
		if err != nil {
			return
		}
		return
	}

	weatherData, err := service.GetWeatherData(city)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get weather data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherJSON, _ := json.Marshal(weatherData)
	redisClient.Set(r.Context(), city, weatherJSON, time.Hour)

	w.Header().Set("Content-Type", "application/json")
	_, err2 := w.Write(weatherJSON)
	if err2 != nil {
		return
	}
}
