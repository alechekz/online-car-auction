package infrastructure

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// MockPricingProvider is a mock implementation of the PricingProvider interface for testing purposes
type MockPricingProvider struct {
	Data *domain.Vehicle
	Err  error
}

// GetRecommendedPrice simulates fetching recommended price for a vehicle
func (m *MockPricingProvider) GetRecommendedPrice(v *domain.Vehicle) (uint64, error) {
	return 99_000, m.Err
}
