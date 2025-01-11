package service

type WeatherResponse struct {
	City  string  `json:"city"`
	Temp  float64 `json:"temperature"`
	Units string  `json:"units"`
}

func GetWeatherData() WeatherResponse {
	return WeatherResponse{
		City:  "Moscow",
		Temp:  -5.0,
		Units: "Celsius",
	}
}
