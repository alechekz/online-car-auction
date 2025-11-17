package usecase

import "github.com/alechekz/online-car-auction/services/inspection/domain"

// BuildDataProvider defines the interface for fetching build data for vehicles
type BuildDataProvider interface {
	Fetch(*domain.Vehicle) error
}
