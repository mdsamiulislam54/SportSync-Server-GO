package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JwtSecret   string
}

func LoadEnv() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		DatabaseURL: os.Getenv("DB_URL"),
		Port:        os.Getenv("PORT"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}

	if cfg.DatabaseURL == "" {
		panic("DB_URL is missing")
	}

	if cfg.JwtSecret == "" {
		panic("JWT_SECRET is missing")
	}

	if cfg.Port == "" {
		cfg.Port = "5000"
	}

	return cfg
}
