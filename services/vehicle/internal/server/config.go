package server

// config holds the configuration for the Vehicle Service server
type config struct {
	Address string
}

// NewConfig creates a new server configuration with default values
func NewConfig(address string) *config {
	return &config{
		Address: address,
	}
}
