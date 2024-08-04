package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Структура для хранения конфигурации приложения

type Config struct {
	PostgresDSN string
	RedisAddr   string
	NatsURL     string
	HTTPPort    string
}

// Загружает переменные окружения из файла .env и создает конфиг

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.\t", err)
	}

	cfg := &Config{
		PostgresDSN: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("SSL_MODE"),
		),
		RedisAddr: fmt.Sprintf("%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		NatsURL: fmt.Sprintf("nats://%s:%s",
			os.Getenv("NATS_HOST"),
			os.Getenv("NATS_PORT"),
		),
		HTTPPort: os.Getenv("HTTP_PORT"),
	}

	return cfg, nil
}
