package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GinMode    string
	HTTPPort   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		GinMode:    GetEnv("GIN_MODE", "debug"),
		HTTPPort:   GetEnv("HTTP_PORT", "8080"),
		DBHost:     GetEnv("DB_HOST", "localhost"),
		DBPort:     GetEnv("DB_PORT", "5432"),
		DBUser:     GetEnv("DB_USER", "user"),
		DBPassword: GetEnv("DB_PASSWORD", "password"),
		DBName:     GetEnv("DB_NAME", "dbname"),
		DBSSLMode:  GetEnv("DB_SSLMODE", "disable"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}