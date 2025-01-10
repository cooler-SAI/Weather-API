package service

import "Weather-API/internal/api"

func GetWeatherData() api.WeatherResponse {
	return api.WeatherResponse{
		City:  "Moscow",
		Temp:  -5.0,
		Units: "Celsius",
	}
}
