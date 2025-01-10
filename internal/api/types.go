package api

type WeatherResponse struct {
	City  string  `json:"city"`
	Temp  float64 `json:"temperature"`
	Units string  `json:"units"`
}
