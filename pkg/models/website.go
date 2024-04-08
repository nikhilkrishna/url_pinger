package models

import (
	"errors"
	"net/url"
)

type WebsiteLog struct {
	SessionId string
	ThreadId int
	URL string 
	Response string
	Error string
	Pattern string  
}


type WebsiteConfig struct {
	URL string
	Pattern string 
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

func WebsiteConfigFactory(url string, pattern string, interval int) (*WebsiteConfig, error) {
	wc := &WebsiteConfig{
		URL: url,
		Pattern: pattern,
		Interval: interval,
	}

	if err := wc.Validate(); err != nil {
		return nil, err
	}

	return wc, nil
}
