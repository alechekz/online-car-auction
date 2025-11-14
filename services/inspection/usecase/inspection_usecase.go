package usecase

import (
	"github.com/alechekz/online-car-auction/services/inspection/domain"
)

// InspectionUsecase defines the interface for inspection-related business logic
type InspectionUsecase interface {
	InspectVehicle(v *domain.Vehicle) error
	GetBuildData(vin string) (*domain.Vehicle, error)
}

// inspectionUsecase is the implementation of InspectionUsecase interface
type inspectionUsecase struct {
	provider BuildDataProvider
}

// NewInspectionUC is the constructor for inspectionUsecase
func NewInspectionUC(p BuildDataProvider) InspectionUsecase {
	return &inspectionUsecase{provider: p}
}

// InspectVehicle inspects a vehicle and creates a new inspection record
func (uc *inspectionUsecase) InspectVehicle(v *domain.Vehicle) error {

	// Validate the inspection data
	if err := v.Validate(); err != nil {
		return domain.ErrValidation
	}
	// Make inspection (dummy logic for example)
	v.Inspect()
	return nil
}

// GetBuildData retrieves the build data for a vehicle by its VIN
func (uc *inspectionUsecase) GetBuildData(vin string) (*domain.Vehicle, error) {

	// Prepare the vehicle instance and validate VIN
	v := &domain.Vehicle{VIN: vin}
	if err := v.ValidateVIN(); err != nil {
		return nil, domain.ErrValidation
	}
	// Fetch build data from the provider
	return v, uc.provider.Fetch(v)
}
