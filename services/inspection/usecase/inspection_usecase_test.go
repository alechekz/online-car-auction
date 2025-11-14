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
	data    func() *domain.Inspection
	isValid bool
}

// newTestInspection is a test valid vehicle instance
func newTestInspection() *domain.Inspection {
	return &domain.Inspection{
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
			data: func() *domain.Inspection {
				return newTestInspection()
			},
			isValid: true,
		},
		{
			name: "invalid VIN",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.VIN = "123"
				return i
			},
			isValid: false,
		},
	}

	// Prepare in-memory repository and usecase
	repo := infrastructure.NewMemoryInspectionRepo()
	provider := infrastructure.NewNHTSABuildDataClient()
	uc := usecase.NewInspectionUC(repo, provider)

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
