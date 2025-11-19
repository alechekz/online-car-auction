package infrastructure

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// MockInspectionProvider is a mock implementation of the InspectionProvider interface for testing purposes
type MockInspectionProvider struct {
	Data *domain.Vehicle
	Err  error
}

// GetBuildData simulates fetching build data for a vehicle
func (m *MockInspectionProvider) GetBuildData(vin string) (*domain.Vehicle, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Data, nil
}

// InspectVehicle simulates vehicle inspection
func (m *MockInspectionProvider) InspectVehicle(v *domain.Vehicle) error {
	return m.Err
}
