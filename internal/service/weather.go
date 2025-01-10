package service

import "weather-api/internal/api"

func GetWeatherData() api.WeatherResponse {
	return api.WeatherResponse{
		City:  "Moscow",
		Temp:  -5.0,
		Units: "Celsius",
	}
}
