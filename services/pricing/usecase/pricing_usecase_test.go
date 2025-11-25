package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction/services/pricing/domain"
	"github.com/alechekz/online-car-auction/services/pricing/infrastructure"
	"github.com/alechekz/online-car-auction/services/pricing/usecase"
)

// test is a struct for vehicle usecase tests
type test struct {
	name    string
	data    func() *domain.Vehicle
	isValid bool
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *domain.Vehicle {
	return &domain.Vehicle{
		VIN:      "1HGCM82633A123456",
		Odometer: 15000,
		Grade:    47,
	}
}

// newTestUC is a helper function to create a PricingUsecase instance for testing
func newTestUC() usecase.PricingUsecase {
	provider := &infrastructure.MockInspectionProvider{
		Data: &domain.Vehicle{
			Msrp: 99_000,
		},
		Err: nil,
	}
	return usecase.NewPricingUC(provider)
}

// TestPricingUsecase_GetRecommendedPrice tests the GetRecommendedPrice method of the PricingUsecase struct
func TestPricingUsecase_GetRecommendedPrice(t *testing.T) {

	// Define test cases
	tests := []test{
		{
			name: "valid vehicle",
			data: func() *domain.Vehicle {
				return newTestVehicle()
			},
			isValid: true,
		},
		{
			name: "invalid VIN",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.VIN = "123"
				return v
			},
			isValid: false,
		},
	}

	uc := newTestUC()

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.GetRecommendedPrice(test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
