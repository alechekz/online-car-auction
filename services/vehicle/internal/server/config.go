package server

import (
	"os"
)

// config holds the configuration for the Vehicle Service server
type config struct {
	Address       string
	InspectionURL string
	DatabaseURL   string
	Repo          string // "postgres" or "inmemory"
}

// NewConfig creates a new server configuration with default values
func NewConfig() *config {
	cfg := &config{
		Address:       ":6061",
		InspectionURL: ":6063",

		Repo: "inmemory",
	}
	if os.Getenv("VEHICLE_URL") != "" {
		cfg.Address = os.Getenv("VEHICLE_URL")
	}
	if os.Getenv("INSPECTION_URL") != "" {
		cfg.InspectionURL = os.Getenv("INSPECTION_URL")
	}
	if os.Getenv("VEHICLE_DB") != "" {
		cfg.DatabaseURL = os.Getenv("VEHICLE_DB")
		cfg.Repo = "postgres"
	}
	return cfg
}
