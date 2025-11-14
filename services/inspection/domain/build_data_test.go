package domain_test

import (
	"testing"

	"github.com/alechekz/online-car-auction/services/inspection/domain"

	"github.com/stretchr/testify/assert"
)

// testBD is a struct for build data tests
type testBD struct {
	name    string
	data    func() *domain.BuildData
	isValid bool
}

// newTestBuildData is a test valid build data instance
func newTestBuildData() *domain.BuildData {
	return &domain.BuildData{
		VIN: "1HGBH41JXMN109186",
	}
}

// TestBuildData_Validate tests the Validate method of the BuildData struct
func TestBuildData_Validate(t *testing.T) {
	tests := []testBD{
		{
			name: "valid build data",
			data: func() *domain.BuildData {
				return newTestBuildData()
			},
			isValid: true,
		},
		{
			name: "missing VIN",
			data: func() *domain.BuildData {
				i := newTestBuildData()
				i.VIN = ""
				return i
			},
			isValid: false,
		},
		{
			name: "invalid VIN",
			data: func() *domain.BuildData {
				i := newTestBuildData()
				i.VIN = "123"
				return i
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
