package unisender

import (
	"time"

	"github.com/basili4-1982/unisender/config"
)

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *config.Config {
	return &config.Config{
		BaseURL:       "https://goapi.unisender.ru/ru/transactional/api/v1",
		Timeout:       30 * time.Second,
		RetryCount:    3,
		RetryWaitTime: 1 * time.Second,
		Debug:         false,
		UserAgent:     "unisender/1.0.0",
	}
}

// Option функция для настройки клиента
type Option func(*config.Config)

// WithBaseURL устанавливает базовый URL API
func WithBaseURL(url string) Option {
	return func(c *config.Config) {
		c.BaseURL = url
	}
}

// WithAPIKey устанавливает API ключ
func WithAPIKey(key string) Option {
	return func(c *config.Config) {
		c.APIKey = key
	}
}

// WithTimeout устанавливает таймаут запросов
func WithTimeout(timeout time.Duration) Option {
	return func(c *config.Config) {
		c.Timeout = timeout
	}
}

// WithRetry настраивает retry механизм
func WithRetry(count int, waitTime time.Duration) Option {
	return func(c *config.Config) {
		c.RetryCount = count
		c.RetryWaitTime = waitTime
	}
}

// WithDebug включает debug режим
func WithDebug(debug bool) Option {
	return func(c *config.Config) {
		c.Debug = debug
	}
}

// WithUserAgent устанавливает User-Agent
func WithUserAgent(ua string) Option {
	return func(c *config.Config) {
		c.UserAgent = ua
	}
}
