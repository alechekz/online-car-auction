package infrastructure

import "github.com/alechekz/online-car-auction/services/pricing/domain"

// MockInspectionProvider is a mock implementation of the InspectionProvider interface for testing purposes
type MockInspectionProvider struct {
	Data *domain.Vehicle
	Err  error
}

// GetMsrp simulates fetching build data for a vehicle
func (m *MockInspectionProvider) GetMsrp(vin string) (uint64, error) {
	if m.Err != nil {
		return 0, m.Err
	}
	return 99_000, nil
}
