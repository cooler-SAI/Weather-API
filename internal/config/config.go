package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found, using local .env")
	}
	return Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
