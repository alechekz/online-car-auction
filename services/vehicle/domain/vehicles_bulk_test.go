package domain_test

import (
	"testing"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"

	"github.com/stretchr/testify/assert"
)

// vBtest is a struct for vehicles bulk tests
type vBtest struct {
	name    string
	data    func() *domain.VehiclesBulk
	isValid bool
}

// newTestVehiclesBulk is a test valid vehicles bulk instance
func newTestVehiclesBulk() *domain.VehiclesBulk {
	return &domain.VehiclesBulk{
		Vehicles: []*domain.Vehicle{newTestVehicle()},
	}
}

// TestVehiclesBulk_Validate tests the Validate method of the VehiclesBulk struct
func TestVehiclesBulk_Validate(t *testing.T) {
	tests := []vBtest{
		{
			name: "valid vehicles bulk",
			data: func() *domain.VehiclesBulk {
				return newTestVehiclesBulk()
			},
			isValid: true,
		},
		{
			name: "empty vehicles slice",
			data: func() *domain.VehiclesBulk {
				vb := newTestVehiclesBulk()
				vb.Vehicles = []*domain.Vehicle{}
				return vb
			},
			isValid: false,
		},
		{
			name: "vehicles slice with invalid vehicle",
			data: func() *domain.VehiclesBulk {
				vb := newTestVehiclesBulk()
				v := newTestVehicle()
				v.VIN = "123"
				vb.Vehicles = append(vb.Vehicles, v)
				return vb
			},
			isValid: false,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.data().Validate()
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
