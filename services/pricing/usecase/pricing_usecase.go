package usecase

import (
	"github.com/alechekz/online-car-auction/services/pricing/domain"
)

// Pricing usecase defines the interface for pricing-related business logic
type PricingUsecase interface {
	GetRecommendedPrice(v *domain.Vehicle) error
}

// pricingUsecase is the implementation of PricingUsecase interface
type pricingUsecase struct {
	provider InspectionProvider
}

// NewPricingUC is the constructor for pricingUsecase
func NewPricingUC(provider InspectionProvider) *pricingUsecase {
	return &pricingUsecase{provider: provider}
}

// GetRecommendedPrice calculates the recommended price for a vehicle
func (uc *pricingUsecase) GetRecommendedPrice(v *domain.Vehicle) error {

	// Validate the vehicle data
	if err := v.Validate(); err != nil {
		return domain.ErrValidation
	}

	// Fetch MRSP
	msrp, err := uc.provider.GetMsrp(v.VIN)
	if err != nil {
		return err
	}
	v.Msrp = msrp

	// Calculate the price
	v.CalcPrice()
	return nil
}
