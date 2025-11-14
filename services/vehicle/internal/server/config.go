package server

import (
	"os"
)

// config holds the configuration for the Vehicle Service server
type config struct {
	Address     string
	DatabaseURL string
	Repo        string // "postgres" or "inmemory"
}

// NewConfig creates a new server configuration with default values
func NewConfig() *config {
	cfg := &config{
		Address: ":6061",

		Repo: "inmemory",
	}
	if os.Getenv("VEHICLE_URL") != "" {
		cfg.Address = os.Getenv("VEHICLE_URL")
	}
	if os.Getenv("VEHICLE_DB") != "" {
		cfg.DatabaseURL = os.Getenv("VEHICLE_DB")
		cfg.Repo = "postgres"
	}
	return cfg
}
