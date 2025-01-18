package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	CacheType string
	RedisAddr string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}
	return Config{
		Port:      getEnv("PORT", "8080"),
		CacheType: getEnv("CACHE_TYPE", "mock"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
