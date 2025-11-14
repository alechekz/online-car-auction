package infrastructure

import (
	"github.com/alechekz/online-car-auction/services/inspection/domain"
)

// MemoryInspectionRepo is an in-memory implementation of InspectionRepository interface
type MemoryInspectionRepo struct {
	data map[string]*domain.Inspection
}

// NewMemoryInspectionRepo creates a new instance of MemoryInspectionRepo
func NewMemoryInspectionRepo() *MemoryInspectionRepo {
	return &MemoryInspectionRepo{
		data: make(map[string]*domain.Inspection),
	}
}

// Save saves a vehicle to the in-memory store
func (r *MemoryInspectionRepo) Save(v *domain.Inspection) error {
	r.data[v.VIN] = v
	return nil
}

// FindByVIN retrieves a vehicle by its VIN
func (r *MemoryInspectionRepo) FindByVIN(vin string) (*domain.Inspection, error) {
	if v, exists := r.data[vin]; exists {
		return v, nil
	}
	return nil, domain.ErrNotFound
}
