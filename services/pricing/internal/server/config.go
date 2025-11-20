package server

import "os"

// config holds the configuration for the Inspection Service server
type config struct {
	HttpAddress   string
	GrpcAddress   string
	DatabaseURL   string
	InspectionURL string
}

// NewConfig creates a new server configuration with default values
func NewConfig() *config {
	cfg := &config{
		HttpAddress:   ":6064",
		GrpcAddress:   ":6065",
		InspectionURL: ":6063",
	}
	if os.Getenv("PRICING_HTTP") != "" {
		cfg.HttpAddress = os.Getenv("PRICING_HTTP")
	}
	if os.Getenv("PRICING_GRPC") != "" {
		cfg.GrpcAddress = os.Getenv("PRICING_GRPC")
	}
	if os.Getenv("INSPECTION_URL") != "" {
		cfg.InspectionURL = os.Getenv("INSPECTION_URL")
	}
	return cfg
}
