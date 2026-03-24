package config

import "time"

type Config struct {
	BaseURL       string
	Timeout       time.Duration
	RetryCount    int
	RetryWaitTime time.Duration
	Debug         bool
	UserAgent     string
	APIKey        string
}
