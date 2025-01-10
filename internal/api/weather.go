package api

import (
	"encoding/json"
	"net/http"
	"weather-api/internal/service"
)

func WeatherHandler(w http.ResponseWriter, _ *http.Request) {
	response := service.GetWeatherData()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
