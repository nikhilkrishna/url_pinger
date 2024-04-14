package config

import (
	"encoding/csv"
	"errors"
	"io"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConn string
}

type WebsiteConfig struct {
	URL      string
	Pattern  string
	Interval int
}

func (wc WebsiteConfig) Validate() error {
	if _, err := url.ParseRequestURI(wc.URL); err != nil {
		return errors.New("invalid URL")
	}

	if wc.Interval <= 0 {
		return errors.New("interval must be positive")
	}

	return nil
}

func NewWebsiteConfig(url string, pattern string, interval int) (*WebsiteConfig, error) {
	wc := &WebsiteConfig{
		URL:      url,
		Pattern:  pattern,
		Interval: interval,
	}

	if err := wc.Validate(); err != nil {
		return nil, err
	}

	return wc, nil
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

func LoadWebsiteSettings(filePath string) ([]*WebsiteConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var configs []*WebsiteConfig

	reader.Read()
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

		wcObj, err := NewWebsiteConfig(record[0], record[1], interval)
		if err != nil {
			return nil, err
		}

		configs = append(configs, wcObj)

	}

	return configs, err
}
