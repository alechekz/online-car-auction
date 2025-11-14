package domain_test

import (
	"testing"

	"github.com/alechekz/online-car-auction/services/inspection/domain"

	"github.com/stretchr/testify/assert"
)

// testI is a struct for inspection tests
type testI struct {
	name     string
	data     func() *domain.Inspection
	isValid  bool
	expected int
}

// newTestInspection is a test valid inspection instance
func newTestInspection() *domain.Inspection {
	return &domain.Inspection{
		VIN:      "1HGBH41JXMN109186",
		Year:     2022,
		Odometer: 12000,
	}
}

// TestInspection_Validate tests the Validate method of the Inspection struct
func TestInspection_Validate(t *testing.T) {
	tests := []testI{
		{
			name: "valid inspection",
			data: func() *domain.Inspection {
				return newTestInspection()
			},
			isValid: true,
		},
		{
			name: "missing VIN",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.VIN = ""
				return i
			},
			isValid: false,
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
		{
			name: "year too old",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.Year = 1800
				return i
			},
			isValid: false,
		},
		{
			name: "year in future",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.Year = 2030
				return i
			},
			isValid: false,
		},
		{
			name: "negative odometer",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.Odometer = -100
				return i
			},
			isValid: false,
		},
		{
			name: "zero odometer",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.Odometer = 0
				return i
			},
			isValid: true,
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

// TestInspection_Inspect tests the Inspect method of the Inspection struct
func TestInspection_Inspect(t *testing.T) {
	tests := []testI{
		{
			name: "only year affects grade",
			data: func() *domain.Inspection {
				return newTestInspection()
			},
			expected: 47,
		},
		{
			name: "year and strong scratches affect grade",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.StrongScratches = true
				return i
			},
			expected: 43,
		},
		{
			name: "year and small scratches affect grade",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.SmallScratches = true
				return i
			},
			expected: 45,
		},
		{
			name: "year and electric fail affect grade",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.ElectricFail = true
				return i
			},
			expected: 43,
		},
		{
			name: "year and suspension fail affect grade",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.SuspensionFail = true
				return i
			},
			expected: 44,
		},
		{
			name: "all factors affect grade",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.StrongScratches = true
				i.SmallScratches = true
				i.ElectricFail = true
				i.SuspensionFail = true
				return i
			},
			expected: 36,
		},
		{
			name: "high odometer affects grade",
			data: func() *domain.Inspection {
				i := newTestInspection()
				i.Odometer = 350000
				return i
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
