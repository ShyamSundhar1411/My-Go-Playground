// Package config loads application configuration from environment
// variables, with sensible defaults for local development.
package config

import "os"

// Config holds all runtime configuration for the API server.
type Config struct {
	Env  string
	Port string
}

// Load reads configuration from the environment. Every field has a
// default so the service runs with zero setup.
func Load() Config {
	return Config{
		Env:  getEnv("APP_ENV", "development"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
