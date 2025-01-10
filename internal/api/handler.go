package api

import (
	"Weather-API/internal/service"
	"encoding/json"
	"net/http"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	response := service.GetWeatherData() // Получение данных о погоде

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
