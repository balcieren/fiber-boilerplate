package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port, DatabaseURI, SigningKey, RedisAddr, RedisPassword string
}

func New() (*Config, error) {
	err := godotenv.Load(".env")
	return &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURI: os.Getenv("DATABASE_URI"),
	}, err
}
