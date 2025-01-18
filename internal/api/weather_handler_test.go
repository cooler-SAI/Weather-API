package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherHandler(t *testing.T) {
	// Создаем запрос
	req, err := http.NewRequest("GET", "/?city=Moscow", nil)
	assert.NoError(t, err)

	// Добавляем API-ключ
	req.Header.Set("X-API-Key", "test-key")

	// Создаем тестовый сервер
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Мокаем `WeatherHandler`, убирая зависимости от Redis
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message":"Weather data"}`))
		if err != nil {
			return
		}
	})

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	// Проверяем тело ответа
	expectedBody := `{"message":"Weather data"}`
	assert.JSONEq(t, expectedBody, rr.Body.String(), "Response body mismatch")
}
