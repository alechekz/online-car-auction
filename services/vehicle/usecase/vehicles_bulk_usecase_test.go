package usecase_test

import (
	"testing"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/infrastructure"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"

	"github.com/stretchr/testify/assert"
)

// newTestVehiclesBulk is a test valid VehiclesBulk instance
func newTestVehiclesBulk() *domain.VehiclesBulk {
	return &domain.VehiclesBulk{
		Vehicles: []*domain.Vehicle{
			newTestVehicle(),
			newTestVehicle(),
			newTestVehicle(),
			newTestVehicle(),
			newTestVehicle(),
			newTestVehicle(),
			newTestVehicle(),
			newTestVehicle(),
		},
	}
}

// TestVehiclesBulkUsecase_Create tests the Create method of VehiclesBulkUsecase
func TestVehiclesBulkUsecase_Create(t *testing.T) {
	mockRepo := new(infrastructure.MockVehiclesRepository)
	vehicleUC := newTestUC()
	uc := usecase.NewVehiclesBulkUC(mockRepo, vehicleUC)
	vb := newTestVehiclesBulk()

	// Successful case
	mockRepo.On("SaveBulk", vb).Return(nil)
	t.Run("valid bulk data", func(t *testing.T) {
		assert.NoError(t, uc.Create(vb))
	})

	// Return error on Validate failure
	mockRepo.On("SaveBulk", vb).Return(nil)
	v := vb.Vehicles[0]
	v.VIN = "123"
	t.Run("invalid bulk data", func(t *testing.T) {
		assert.Error(t, uc.Create(vb))
	})

	// Return error on repository failure
	mockRepo.On("SaveBulk", vb).Return(assert.AnError)
	t.Run("repository failure", func(t *testing.T) {
		assert.Error(t, uc.Create(vb))
	})
}

// TestVehiclesBulkUsecase_Update tests the Update method of VehiclesBulkUsecase
func TestVehiclesBulkUsecase_Update(t *testing.T) {
	mockRepo := new(infrastructure.MockVehiclesRepository)
	vehicleUC := newTestUC()
	uc := usecase.NewVehiclesBulkUC(mockRepo, vehicleUC)
	vb := newTestVehiclesBulk()

	// Successful case
	mockRepo.On("UpdateBulk", vb).Return(nil)
	t.Run("valid bulk data", func(t *testing.T) {
		assert.NoError(t, uc.Update(vb))
	})

	// Return error on Validate failure
	mockRepo.On("UpdateBulk", vb).Return(nil)
	v := vb.Vehicles[0]
	v.VIN = "123"
	t.Run("invalid bulk data", func(t *testing.T) {
		assert.Error(t, uc.Update(vb))
	})

	// Return error on repository failure
	mockRepo.On("UpdateBulk", vb).Return(assert.AnError)
	t.Run("repository failure", func(t *testing.T) {
		assert.Error(t, uc.Update(vb))
	})
}
