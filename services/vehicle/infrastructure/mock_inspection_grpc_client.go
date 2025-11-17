package infrastructure

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// MockBuildDataProvider is a mock implementation of the BuildDataProvider interface for testing purposes
type MockBuildDataProvider struct {
	Data *domain.Vehicle
	Err  error
}

// Ensure MockBuildDataProvider implements BuildDataProvider interface
func (m *MockBuildDataProvider) Fetch(vin string) (*domain.Vehicle, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Data, nil
}
