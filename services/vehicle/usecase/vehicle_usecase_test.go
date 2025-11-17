package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/infrastructure"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
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
		Year:     2020,
		Odometer: 15000,
		MSRP:     25000.00,
	}
}

// TestVehicleUsecase_CreateVehicle tests the CreateVehicle method of the VehicleUsecase struct
func TestVehicleUsecase_CreateVehicle(t *testing.T) {

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

	// Prepare in-memory repository and usecase
	repo := infrastructure.NewMemoryVehicleRepo()
	provider := &infrastructure.InspectionGRPCClient{} // You might want to use a mock here
	uc := usecase.NewVehicleUC(repo, provider)

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.CreateVehicle(test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// TestVehicleUsecase_GetVehicle tests the GetVehicle method of the VehicleUsecase struct
func TestVehicleUsecase_GetVehicle(t *testing.T) {

	// Prepare in-memory repository and usecase
	repo := infrastructure.NewMemoryVehicleRepo()
	uc := usecase.NewVehicleUC(repo)
	v := newTestVehicle()
	err := uc.CreateVehicle(v)
	assert.NoError(t, err)

	// Valid case
	t.Run("existing vehicle", func(t *testing.T) {
		got, err := uc.GetVehicle(v.VIN)
		assert.NoError(t, err)
		assert.Equal(t, v.VIN, got.VIN)
	})

	// Invalid case
	t.Run("non-existing vehicle", func(t *testing.T) {
		_, err := uc.GetVehicle("NONEXISTENTVIN12345")
		assert.Error(t, err)
	})
}

// TestVehicleUsecase_UpdateVehicle tests the UpdateVehicle method of the VehicleUsecase struct
func TestVehicleUsecase_UpdateVehicle(t *testing.T) {

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
			name: "invalid vehicle, too old year",
			data: func() *domain.Vehicle {
				v := newTestVehicle()
				v.Year = 1800
				return v
			},
			isValid: false,
		},
	}

	// Prepare in-memory repository and usecase
	repo := infrastructure.NewMemoryVehicleRepo()
	uc := usecase.NewVehicleUC(repo)
	v := newTestVehicle()
	err := uc.CreateVehicle(v)
	assert.NoError(t, err)

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := uc.UpdateVehicle(test.data())
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// TestVehicleUsecase_DeleteVehicle tests the DeleteVehicle method of the VehicleUsecase struct
func TestVehicleUsecase_DeleteVehicle(t *testing.T) {

	// Prepare in-memory repository and usecase
	repo := infrastructure.NewMemoryVehicleRepo()
	uc := usecase.NewVehicleUC(repo)
	v := newTestVehicle()
	err := uc.CreateVehicle(v)
	assert.NoError(t, err)

	// Valid case
	t.Run("delete existing vehicle", func(t *testing.T) {
		err := uc.DeleteVehicle(v.VIN)
		assert.NoError(t, err)

		// Verify deletion
		_, err = uc.GetVehicle(v.VIN)
		assert.Error(t, err)
	})

	// Invalid case
	t.Run("delete non-existing vehicle", func(t *testing.T) {
		err := uc.DeleteVehicle("NONEXISTENTVIN12345")
		assert.Error(t, err)
	})
}

// TestVehicleUsecase_ListVehicles tests the ListVehicles method of the VehicleUsecase struct
func TestVehicleUsecase_ListVehicles(t *testing.T) {

	// Prepare in-memory repository and usecase
	repo := infrastructure.NewMemoryVehicleRepo()
	uc := usecase.NewVehicleUC(repo)

	// Empty list case
	t.Run("empty list", func(t *testing.T) {
		vehicles, err := uc.ListVehicles()
		assert.NoError(t, err)
		assert.Len(t, vehicles, 0)
	})

	// List with vehicles case
	t.Run("list with multiple vehicles", func(t *testing.T) {
		v1 := newTestVehicle()
		v2 := &domain.Vehicle{
			VIN:      "2HGCM82633A654321",
			Year:     2021,
			Odometer: 5000,
			MSRP:     30000,
		}
		assert.NoError(t, uc.CreateVehicle(v1))
		assert.NoError(t, uc.CreateVehicle(v2))

		// Retrieve and verify list
		vehicles, err := uc.ListVehicles()
		assert.NoError(t, err)
		assert.Len(t, vehicles, 2)

		// Verify that both vehicles are in the list
		vins := []string{vehicles[0].VIN, vehicles[1].VIN}
		assert.Contains(t, vins, v1.VIN)
		assert.Contains(t, vins, v2.VIN)
	})
}
