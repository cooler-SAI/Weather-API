package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
	"weather-api/internal/service"

	"github.com/rs/zerolog/log"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

type MockCache struct {
	data map[string]string
}

func NewMockCache() *MockCache {
	return &MockCache{data: make(map[string]string)}
}

func (m *MockCache) Get(key string) (string, error) {
	if val, ok := m.data[key]; ok {
		return val, nil
	}
	return "", errors.New("cache miss")
}

func (m *MockCache) Set(key string, value interface{}, _ time.Duration) error {
	if strValue, ok := value.([]byte); ok {
		m.data[key] = string(strValue)
	} else {
		m.data[key] = value.(string)
	}
	return nil
}

var cacheClient Cache = NewMockCache()

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	expectedKey := os.Getenv("OPENWEATHER_API_KEY")

	if apiKey != expectedKey {
		log.Warn().Msg("Proceeding with default data, no valid API key provided")
		defaultData := map[string]interface{}{
			"city":  "Unknown",
			"temp":  0,
			"units": "Celsius",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(defaultData)
		if err != nil {
			return
		}
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
		_, err := w.Write([]byte(cachedData))
		if err != nil {
			return
		}
		return
	}

	weatherData, err := service.GetWeatherData(city)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch weather data")
		http.Error(w, `{"error":"Failed to fetch weather data"}`, http.StatusInternalServerError)
		return
	}

	weatherJSON, _ := json.Marshal(weatherData)

	err2 := cacheClient.Set(city, weatherJSON, time.Hour)
	if err2 != nil {
		log.Warn().Err(err2).Msg("Failed to cache weather data")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err3 := w.Write(weatherJSON)
	if err3 != nil {
		return
	}
}
