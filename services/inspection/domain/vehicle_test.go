package domain_test

import (
	"testing"

	"github.com/alechekz/online-car-auction/services/inspection/domain"

	"github.com/stretchr/testify/assert"
)

// test is a struct for build data tests
type test struct {
	name     string
	data     func() *domain.Vehicle
	isValid  bool
	expected int
}

// newTestVehicle is a test valid vehicle instance
func newTestVehicle() *domain.Vehicle {
	return &domain.Vehicle{
		VIN:      "1HGBH41JXMN109186",
		Year:     2022,
		Odometer: 12000,
	}
}

// TestVehicle_ValidateVIN tests the ValidateVIN method of the Vehicle struct
func TestVehicle_ValidateVIN(t *testing.T) {
	tests := []test{
		{
			name: "valid VIN",
			data: func() *domain.Vehicle {
				return newTestVehicle()
			},
			isValid: true,
		},
		{
			name: "missing VIN",
			data: func() *domain.Vehicle {
				i := newTestVehicle()
				i.VIN = ""
				return i
			},
			isValid: false,
		},
		{
			name: "invalid VIN",
			data: func() *domain.Vehicle {
				i := newTestVehicle()
				i.VIN = "123"
				return i
			},
			isValid: false,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.data().ValidateVIN()
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
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
			name: "missing VIN",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.VIN = ""
				return v
			},
			isValid: false,
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
			name: "year too old",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Year = 1800
				return v
			},
			isValid: false,
		},
		{
			name: "year in future",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Year = 2030
				return v
			},
			isValid: false,
		},
		{
			name: "negative odometer",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Odometer = -100
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

// TestVehicle_Inspect tests the Inspect method of the Vehicle struct
func TestVehicle_Inspect(t *testing.T) {
	tests := []test{
		{
			name: "only year affects grade",
			data: func() *domain.Vehicle {
				return newTestVehicle()
			},
			expected: 47,
		},
		{
			name: "year and strong scratches affect grade",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.StrongScratches = true
				return v
			},
			expected: 43,
		},
		{
			name: "year and small scratches affect grade",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.SmallScratches = true
				return v
			},
			expected: 45,
		},
		{
			name: "year and electric fail affect grade",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.ElectricFail = true
				return v
			},
			expected: 43,
		},
		{
			name: "year and suspension fail affect grade",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.SuspensionFail = true
				return v
			},
			expected: 44,
		},
		{
			name: "all factors affect grade",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.StrongScratches = true
				v.SmallScratches = true
				v.ElectricFail = true
				v.SuspensionFail = true
				return v
			},
			expected: 36,
		},
		{
			name: "high odometer affects grade",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Odometer = 350000
				return v
			},
			expected: 30,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inspection := test.data()
			inspection.Inspect()
			assert.Equal(t, test.expected, inspection.Grade)
		})
	}
}
