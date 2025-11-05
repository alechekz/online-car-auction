package repository

import "github.com/alechekz/online-car-auction/services/vehicle/domain"

// VehicleRepository defines the interface for vehicle data operations
type VehicleRepository interface {
	Save(v *domain.Vehicle) error
	FindByVIN(vin string) (*domain.Vehicle, error)
	Update(v *domain.Vehicle) error
	Delete(vin string) error
	List() ([]*domain.Vehicle, error)
}
