package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction/services/inspection/domain"
	"github.com/alechekz/online-car-auction/services/inspection/infrastructure"
	"github.com/alechekz/online-car-auction/services/inspection/usecase"
)

// test is a struct for inspection usecase tests
type test struct {
	name    string
	data    func() *domain.Vehicle
	isValid bool
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *domain.Vehicle {
	return &domain.Vehicle{
		VIN:      "1HGCM82633A123456",
		Year:     2022,
		Odometer: 15000,
	}
}

// TestInspectionUsecase_InspectVehicle tests the InspectVehicle method of the InspectionUsecase struct
func TestInspectionUsecase_InspectVehicle(t *testing.T) {

	// Define test cases
	tests := []test{
		{
			name: "valid inspection",
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

	// Prepare in-memory repository and usecase
	provider := infrastructure.NewNHTSABuildDataClient()
	msrp := infrastructure.NewMockMSRPClient()
	uc := usecase.NewInspectionUC(provider, msrp)

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.InspectVehicle(test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
