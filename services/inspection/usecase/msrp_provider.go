package usecase

import "github.com/alechekz/online-car-auction/services/inspection/domain"

// MSRPProvider defines the interface for fetching MSRP for vehicles
type MSRPProvider interface {
	Fetch(v *domain.Vehicle) error
}
