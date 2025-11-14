package server

import "os"

// config holds the configuration for the Inspection Service server
type config struct {
	Address     string
	DatabaseURL string
}

// NewConfig creates a new server configuration with default values
func NewConfig() *config {
	cfg := &config{
		Address: ":6062",
	}
	if os.Getenv("INSPECTION_URL") != "" {
		cfg.Address = os.Getenv("INSPECTION_URL")
	}
	return cfg
}
