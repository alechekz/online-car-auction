package domain_test

import (
	"testing"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"

	"github.com/stretchr/testify/assert"
)

// test is a struct for vehicle tests
type test struct {
	name    string
	data    func() *domain.Vehicle
	isValid bool
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *domain.Vehicle {
	return &domain.Vehicle{
		VIN:      "1HGBH41JXMN109186",
		Year:     2022,
		MSRP:     25999.99,
		Odometer: 12000,
	}
}

// TestVehicle_Validate tests the Validate method of the Vehicle struct
func TestVehicle_Validate(t *testing.T) {
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
		{
			name: "too old year",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Year = 1800
				return v
			},
			isValid: false,
		},
		{
			name: "too new year",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Year = 2026
				return v
			},
			isValid: false,
		},
		{
			name: "zero MSRP",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.MSRP = 0.0
				return v
			},
			isValid: false,
		},
		{
			name: "negative MSRP",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.MSRP = -1000.00
				return v
			},
			isValid: false,
		},
		{
			name: "zero odometer",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Odometer = 0
				return v
			},
			isValid: true,
		},
		{
			name: "negative odometer",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Odometer = -5000
				return v
			},
			isValid: false,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid {
				assert.NoError(t, test.data().Validate())
			} else {
				assert.Error(t, test.data().Validate())
			}
		})
	}
}
