package usecase

import (
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/repository"
)

// VehicleUsecase defines the interface for vehicle-related business logic
type VehicleUsecase interface {
	Create(v *domain.Vehicle) error
	Get(vin string) (*domain.Vehicle, error)
	Update(v *domain.Vehicle) error
	Delete(vin string) error
	List() ([]*domain.Vehicle, error)
	Fetch(v *domain.Vehicle) error
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

// Fetch fetches all necessary data for vehicle processing
func (uc *vehicleUsecase) Fetch(v *domain.Vehicle) error {

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

// Create creates a new vehicle record
func (uc *vehicleUsecase) Create(v *domain.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return domain.ErrValidation
	}

	// Fetch all necessary data for vehicle processing
	if err := uc.Fetch(v); err != nil {
		return err
	}

	// Save the vehicle record
	return uc.repo.Save(v)
}

// Get retrieves a vehicle by its VIN
func (uc *vehicleUsecase) Get(vin string) (*domain.Vehicle, error) {
	v, err := uc.repo.FindByVIN(vin)
	if err != nil {
		return nil, domain.ErrNotFound
	}
	if v == nil {
		return nil, domain.ErrNotFound
	}
	return v, nil
}

// Update updates an existing vehicle record
func (uc *vehicleUsecase) Update(v *domain.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return domain.ErrValidation
	}

	// Fetch all necessary data for vehicle processing
	if err := uc.Fetch(v); err != nil {
		return err
	}

	// Update the vehicle record
	if err := uc.repo.Update(v); err != nil {
		return domain.ErrNotFound
	}
	return nil
}

// Delete deletes a vehicle by its VIN
func (uc *vehicleUsecase) Delete(vin string) error {
	if err := uc.repo.Delete(vin); err != nil {
		return domain.ErrNotFound
	}
	return nil
}

// List lists all vehicles
func (uc *vehicleUsecase) List() ([]*domain.Vehicle, error) {
	return uc.repo.List()
}
