package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherHandler(t *testing.T) {
	// Создаем запрос
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Создаем тестовый сервер
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WeatherHandler)

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	// Проверяем содержимое ответа
	expectedBody := `{"message":"Weather data"}` // Убедись, что это корректный ответ
	assert.JSONEq(t, expectedBody, rr.Body.String(), "Response body mismatch")
}
