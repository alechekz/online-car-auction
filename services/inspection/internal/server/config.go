package server

import "os"

// config holds the configuration for the Inspection Service server
type config struct {
	HttpAddress string
	GrpcAddress string
	DatabaseURL string
}

// NewConfig creates a new server configuration with default values
func NewConfig() *config {
	cfg := &config{
		HttpAddress: ":6062",
		GrpcAddress: ":6063",
	}
	if os.Getenv("INSPECTION_HTTP") != "" {
		cfg.HttpAddress = os.Getenv("INSPECTION_HTTP")
	}
	if os.Getenv("INSPECTION_GRPC") != "" {
		cfg.GrpcAddress = os.Getenv("INSPECTION_GRPC")
	}
	return cfg
}
