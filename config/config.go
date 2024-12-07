package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, relying on environment variables")
	}

	dbURL := os.Getenv("DB_CONN_STRING")
	serverPort := os.Getenv("SERVER_PORT")

	if dbURL == "" || serverPort == "" {
		log.Fatal("Environment variables DB_CONN_STRING or SERVER_PORT are not set")
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  serverPort,
	}
}
