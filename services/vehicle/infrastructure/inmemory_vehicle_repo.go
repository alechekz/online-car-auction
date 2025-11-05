package infrastructure

import (
	"errors"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"
)

// MemoryVehicleRepo is an in-memory implementation of VehicleRepository interface
type MemoryVehicleRepo struct {
	data map[string]*domain.Vehicle
}

// NewMemoryVehicleRepo creates a new instance of MemoryVehicleRepo
func NewMemoryVehicleRepo() *MemoryVehicleRepo {
	return &MemoryVehicleRepo{
		data: make(map[string]*domain.Vehicle),
	}
}

// Save saves a vehicle to the in-memory store
func (r *MemoryVehicleRepo) Save(v *domain.Vehicle) error {
	r.data[v.VIN] = v
	return nil
}

// FindByVIN retrieves a vehicle by its VIN
func (r *MemoryVehicleRepo) FindByVIN(vin string) (*domain.Vehicle, error) {
	if v, ok := r.data[vin]; ok {
		return v, nil
	}
	return nil, errors.New("not found")
}

// Update updates an existing vehicle in the in-memory store
func (r *MemoryVehicleRepo) Update(v *domain.Vehicle) error {
	if _, ok := r.data[v.VIN]; !ok {
		return errors.New("not found")
	}
	r.data[v.VIN] = v
	return nil
}

// Delete removes a vehicle from the in-memory store by its VIN
func (r *MemoryVehicleRepo) Delete(vin string) error {
	if _, ok := r.data[vin]; !ok {
		return errors.New("vehicle not found")
	}
	delete(r.data, vin)
	return nil
}

// List lists all vehicles in the in-memory store
func (r *MemoryVehicleRepo) List() ([]*domain.Vehicle, error) {
	result := make([]*domain.Vehicle, 0, len(r.data))
	for _, v := range r.data {
		result = append(result, v)
	}
	return result, nil
}
