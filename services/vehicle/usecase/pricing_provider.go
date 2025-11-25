package usecase

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// PricingProvider defines the interface for fetching vehicle pricing data
type PricingProvider interface {
	GetRecommendedPrice(v *domain.Vehicle) (uint64, error)
}
