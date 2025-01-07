package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds the application configuration.
type Config struct {
	UserRequestLimit   int // Per-user limit
	GlobalRequestLimit int // Global traffic limit
}

// LoadConfig loads configuration values from environment variables.
func LoadConfig() *Config {
	// Load per-user request limit
	userLimitStr := os.Getenv("USER_REQUEST_LIMIT")
	userLimit, err := strconv.Atoi(userLimitStr)
	if err != nil {
		log.Fatalf("Invalid USER_REQUEST_LIMIT value: %v", err)
	}

	// Load global traffic limit
	globalLimitStr := os.Getenv("GLOBAL_REQUEST_LIMIT")
	globalLimit, err := strconv.Atoi(globalLimitStr)
	if err != nil {
		log.Fatalf("Invalid GLOBAL_REQUEST_LIMIT value: %v", err)
	}

	return &Config{
		UserRequestLimit:   userLimit,
		GlobalRequestLimit: globalLimit,
	}
}
