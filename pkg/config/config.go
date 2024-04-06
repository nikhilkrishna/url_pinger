package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBConn string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{
		DBConn: os.Getenv("DB_CONN"),
	}, nil
}