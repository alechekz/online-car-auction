package usecase

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// InspectionProvider defines the interface for fetching vehicle inspection data
type InspectionProvider interface {
	InspectVehicle(v *domain.Vehicle) error
	GetBuildData(vin string) (*domain.Vehicle, error)
}
