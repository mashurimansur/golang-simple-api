package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database string
	Port     string
	Host     string
	Username string
	Password string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Database: os.Getenv("DB_DATABASE"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
