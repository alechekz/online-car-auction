package usecase

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// BuildDataProvider defines the interface for fetching build data for vehicles
type BuildDataProvider interface {
	Fetch(vin string) (*domain.Vehicle, error)
}
