package config

import (
	"fmt"
	"os"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	Host       string
	Port       string
	LogLevel   string
	SentryDSN  string
	SentryEnv  string
	AppVersion string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() (*Config, error) {
	cfg := &Config{
		Host:       envOrDefault("HOST", "0.0.0.0"),
		Port:       envOrDefault("PORT", "8080"),
		LogLevel:   envOrDefault("LOG_LEVEL", "info"),
		SentryDSN:  envOrDefault("SENTRY_DSN", ""),
		SentryEnv:  envOrDefault("SENTRY_ENVIRONMENT", "development"),
		AppVersion: envOrDefault("APP_VERSION", "0.1.0"),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// ListenAddr returns the combined host:port string for the HTTP listener.
func (c *Config) ListenAddr() string {
	return c.Host + ":" + c.Port
}

// SentryEnabled returns true if a Sentry DSN has been configured.
func (c *Config) SentryEnabled() bool {
	return c.SentryDSN != ""
}

func (c *Config) validate() error {
	if c.Port == "" {
		return fmt.Errorf("PORT must not be empty")
	}
	return nil
}

func envOrDefault(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
