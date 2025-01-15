package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	City  string  `json:"city"`
	Temp  float64 `json:"temperature"`
	Units string  `json:"units"`
}

type openWeatherMapResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

const (
	apiKey = "2bfcf1d1a1c828a2562683b20b52eaa7"
)

func GetWeatherData(city string) (WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var owmResp openWeatherMapResponse
	if err := json.NewDecoder(resp.Body).Decode(&owmResp); err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return WeatherResponse{
		City:  owmResp.Name,
		Temp:  owmResp.Main.Temp,
		Units: "Celsius",
	}, nil
}
