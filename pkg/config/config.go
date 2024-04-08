package config

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"url_pinger/pkg/models"

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


func LoadWebsiteSettings(filePath string) ([]*models.WebsiteConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var configs []*models.WebsiteConfig

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break 
			}
			return nil, err
		}

		interval, err := strconv.Atoi((record[2]))
		if err != nil {
			return nil, err
		}

		wcObj, err := models.WebsiteConfigFactory(record[0],record[1],interval)
		if err != nil {
			return nil, err
		}

		configs = append(configs, wcObj)

	}

	return configs, err
}