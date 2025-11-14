package server

import (
	"os"
)

// config holds the configuration for the Inspection Service server
type config struct {
	Address     string
	DatabaseURL string
	Repo        string // "postgres" or "inmemory"
}

// NewConfig creates a new server configuration with default values
func NewConfig() *config {
	cfg := &config{
		Address: ":6062",
		Repo:    "inmemory",
	}
	if os.Getenv("INSPECTION_URL") != "" {
		cfg.Address = os.Getenv("INSPECTION_URL")
	}
	if os.Getenv("INSPECTION_DB") != "" {
		cfg.DatabaseURL = os.Getenv("INSPECTION_DB")
		cfg.Repo = "postgres"
	}
	return cfg
}
