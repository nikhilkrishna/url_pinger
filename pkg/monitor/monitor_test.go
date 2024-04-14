package monitor

import (
	"testing"
	"url_pinger/pkg/config"

	"github.com/stretchr/testify/assert"
)

// Creates a new WebsiteConfig object with the provided URL, pattern, and interval
func TestWebsiteConfigFactoryValidInput(t *testing.T) {
	url := "https://example.com"
	pattern := "example"
	interval := 5

	wc, err := config.NewWebsiteConfig(url, pattern, interval)

	assert.Nil(t, err)
	assert.Equal(t, url, wc.URL)
	assert.Equal(t, pattern, wc.Pattern)
	assert.Equal(t, interval, wc.Interval)
}

// Returns an error if the provided URL is invalid
func TestWebsiteConfigFactoryInvalidURL(t *testing.T) {
	url := "invalid-url"
	pattern := "example"
	interval := 5

	wc, err := config.NewWebsiteConfig(url, pattern, interval)

	assert.NotNil(t, err)
	assert.Nil(t, wc)
}

// Returns an error if the provided interval is less than or equal to zero
func TestWebsiteConfigFactoryInvalidInterval(t *testing.T) {
	url := "https://example.com"
	pattern := "example"
	interval := 0

	wc, err := config.NewWebsiteConfig(url, pattern, interval)

	assert.NotNil(t, err)
	assert.Nil(t, wc)
}
