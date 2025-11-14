package usecase

import (
	"github.com/alechekz/online-car-auction/services/inspection/domain"
	"github.com/alechekz/online-car-auction/services/inspection/repository"
)

// InspectionUsecase defines the interface for inspection-related business logic
type InspectionUsecase interface {
	InspectVehicle(v *domain.Inspection) error
	GetBuildData(vin string) (*domain.BuildData, error)
}

// inspectionUsecase is the implementation of InspectionUsecase interface
type inspectionUsecase struct {
	repo     repository.InspectionRepository
	provider BuildDataProvider
}

// NewInspectionUC is the constructor for inspectionUsecase
func NewInspectionUC(r repository.InspectionRepository, p BuildDataProvider) InspectionUsecase {
	return &inspectionUsecase{repo: r, provider: p}
}

// InspectVehicle inspects a vehicle and creates a new inspection record
func (uc *inspectionUsecase) InspectVehicle(i *domain.Inspection) error {

	// Validate the inspection data
	if err := i.Validate(); err != nil {
		return domain.ErrValidation
	}
	// Make inspection (dummy logic for example)
	i.Inspect()

	// Save the inspection record
	return uc.repo.Save(i)
}

// GetBuildData retrieves the build data for a vehicle by its VIN
func (uc *inspectionUsecase) GetBuildData(vin string) (*domain.BuildData, error) {

	// Prepare build data struct and validate VIN
	buildData := &domain.BuildData{VIN: vin}
	if err := buildData.Validate(); err != nil {
		return nil, domain.ErrValidation
	}

	// Fetch build data from the provider
	return buildData, uc.provider.Fetch(buildData)
}
