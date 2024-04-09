package tests

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"url_pinger/pkg/config"
	"github.com/stretchr/testify/require"
)


func TestLoadConfig(t *testing.T) {
	path := filepath.Join("test.env")
	cfg, err := config.LoadConfig(path)

	require.NoError(t, err)
    require.NotEmpty(t, cfg.DBConn)

}
	
func TestLoadWebsiteSettings_ValidCSVFile(t *testing.T) {
	// Create a temporary CSV file with valid data
	file, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write valid data to the CSV file
	data := []string{"https://www.example.com", "example", "60"}
	writer := csv.NewWriter(file)
	writer.Write(data)
	writer.Flush()


	// Load website settings from the CSV file
	configs, err := config.LoadWebsiteSettings(file.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert that the correct number of website configs are loaded
	if len(configs) != 1 {
		t.Errorf("Expected 1 website config, got %d", len(configs))
	}

	// Assert that the loaded website config has the correct values
	expectedConfig := &config.WebsiteConfig{
		URL:      "https://www.example.com",
		Pattern:  "example",
		Interval: 60,
	}
	if !reflect.DeepEqual(configs[0], expectedConfig) {
		t.Errorf("Expected website config %+v, got %+v", expectedConfig, configs[0])
	}
}

	// Returns an error for an invalid CSV file path
func TestLoadWebsiteSettings_InvalidCSVFilePath(t *testing.T) {
	// Load website settings from an invalid CSV file path
	_, err := config.LoadWebsiteSettings("invalid.csv")

	// Assert that an error is returned
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
