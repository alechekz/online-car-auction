package usecase

import (
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/repository"
)

// VehicleUsecase defines the interface for vehicle-related business logic
type VehicleUsecase interface {
	CreateVehicle(v *domain.Vehicle) error
	GetVehicle(vin string) (*domain.Vehicle, error)
	UpdateVehicle(v *domain.Vehicle) error
	DeleteVehicle(vin string) error
	ListVehicles() ([]*domain.Vehicle, error)
}

// vehicleUsecase is the implementation of VehicleUsecase interface
type vehicleUsecase struct {
	repo               repository.VehicleRepository
	inspectionProvider InspectionProvider
	pricingProvider    PricingProvider
}

// NewVehicleUC is the constructor for vehicleUsecase
func NewVehicleUC(r repository.VehicleRepository, inspectionProvider InspectionProvider, pricingProvider PricingProvider) *vehicleUsecase {
	return &vehicleUsecase{
		repo:               r,
		inspectionProvider: inspectionProvider,
		pricingProvider:    pricingProvider,
	}
}

// fetch fetches all necessary data for vehicle processing
func (uc *vehicleUsecase) fetch(v *domain.Vehicle) error {

	// Fetch build data and merge with user's vehicle data
	bd, err := uc.inspectionProvider.GetBuildData(v.VIN)
	if err != nil {
		return err
	}
	if v.Brand == "" {
		v.Brand = bd.Brand
	}
	if v.Engine == "" {
		v.Engine = bd.Engine
	}
	if v.Transmission == "" {
		v.Transmission = bd.Transmission
	}
	v.MSRP = bd.MSRP

	// Perform vehicle inspection to get the grade
	if err := uc.inspectionProvider.InspectVehicle(v); err != nil {
		return err
	}

	// Calculate the price based on MSRP and grade
	price, err := uc.pricingProvider.GetRecommendedPrice(v)
	if err != nil {
		return err
	}
	v.Price = price
	return nil
}

// CreateVehicle creates a new vehicle record
func (uc *vehicleUsecase) CreateVehicle(v *domain.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return domain.ErrValidation
	}

	// Fetch all necessary data for vehicle processing
	if err := uc.fetch(v); err != nil {
		return err
	}

	// Save the vehicle record
	return uc.repo.Save(v)
}

// GetVehicle retrieves a vehicle by its VIN
func (uc *vehicleUsecase) GetVehicle(vin string) (*domain.Vehicle, error) {
	v, err := uc.repo.FindByVIN(vin)
	if err != nil {
		return nil, domain.ErrNotFound
	}
	if v == nil {
		return nil, domain.ErrNotFound
	}
	return v, nil
}

// UpdateVehicle updates an existing vehicle record
func (uc *vehicleUsecase) UpdateVehicle(v *domain.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return domain.ErrValidation
	}

	// Fetch all necessary data for vehicle processing
	if err := uc.fetch(v); err != nil {
		return err
	}

	// Update the vehicle record
	if err := uc.repo.Update(v); err != nil {
		return domain.ErrNotFound
	}
	return nil
}

// DeleteVehicle deletes a vehicle by its VIN
func (uc *vehicleUsecase) DeleteVehicle(vin string) error {
	if err := uc.repo.Delete(vin); err != nil {
		return domain.ErrNotFound
	}
	return nil
}

// ListVehicles lists all vehicles
func (uc *vehicleUsecase) ListVehicles() ([]*domain.Vehicle, error) {
	return uc.repo.List()
}
