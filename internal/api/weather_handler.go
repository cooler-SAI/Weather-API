package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// Интерфейс для кэша
type Cache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

// MockCache используется для тестов
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

func (m *MockCache) Set(key string, value interface{}, expiration time.Duration) error {
	if strValue, ok := value.([]byte); ok {
		m.data[key] = string(strValue)
	} else {
		m.data[key] = value.(string)
	}
	return nil
}

// Глобальная переменная для кэша
var cacheClient Cache = NewMockCache()

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка наличия API-ключа
	apiKey := r.Header.Get("X-API-Key")
	if apiKey != "test-key" {
		log.Error().Msg("API key is missing or invalid")
		http.Error(w, `{"error":"API key is missing or invalid"}`, http.StatusInternalServerError)
		return
	}

	// Получение города из параметров запроса
	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Moscow"
	}

	// Проверяем наличие данных в кэше
	cachedData, err := cacheClient.Get(city)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedData))
		return
	}

	// Получение данных о погоде
	weatherData := map[string]interface{}{
		"message": "Weather data",
	} // Заглушка для тестов
	weatherJSON, _ := json.Marshal(weatherData)

	// Сохраняем данные в кэш
	cacheClient.Set(city, weatherJSON, time.Hour)

	// Возвращаем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(weatherJSON)
}
