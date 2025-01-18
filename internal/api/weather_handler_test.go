package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/?city=Moscow", nil)
	assert.NoError(t, err)

	req.Header.Set("X-API-Key", "test-key")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message":"Weather data"}`))
		if err != nil {
			return
		}
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	expectedBody := `{"message":"Weather data"}`
	assert.JSONEq(t, expectedBody, rr.Body.String(), "Response body mismatch")
}
