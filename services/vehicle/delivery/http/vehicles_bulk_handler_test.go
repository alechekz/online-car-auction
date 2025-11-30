package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"
)

// TestVehicleHandler_CreateVehiclesBulk tests the CreateVehiclesBulk HTTP handler
func TestVehicleHandler_CreateVehiclesBulk(t *testing.T) {
	router := NewTestRouter()

	// Valid case
	t.Run("valid request", func(t *testing.T) {
		vb := domain.VehiclesBulk{
			Vehicles: []*domain.Vehicle{
				{VIN: "1HGCM82633A123456", Year: 2020, Odometer: 15000},
				{VIN: "2HGCM82633A654321", Year: 2019, Odometer: 30000},
			},
		}
		body, _ := json.Marshal(vb)

		req := httptest.NewRequest(http.MethodPost, "/vehicles/bulk", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		var got domain.VehiclesBulk
		err := json.NewDecoder(rec.Body).Decode(&got)
		assert.NoError(t, err)
		assert.Equal(t, vb.Vehicles[0].VIN, got.Vehicles[0].VIN)
	})

	// Invalid case
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/vehicles/bulk", bytes.NewReader([]byte(`{"vehicles":"oops"}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

// TestVehicleHandler_UpdateVehiclesBulk tests the UpdateVehiclesBulk HTTP handler
func TestVehicleHandler_UpdateVehiclesBulk(t *testing.T) {
	router := NewTestRouter()

	// Create vehicles first
	vb := domain.VehiclesBulk{
		Vehicles: []*domain.Vehicle{
			{VIN: "1HGCM82633A123456", Year: 2020, Odometer: 15000},
			{VIN: "2HGCM82633A654321", Year: 2019, Odometer: 30000},
		},
	}
	body, _ := json.Marshal(vb)
	req := httptest.NewRequest(http.MethodPost, "/vehicles/bulk", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Valid case
	t.Run("valid request", func(t *testing.T) {

		vb := domain.VehiclesBulk{
			Vehicles: []*domain.Vehicle{
				{VIN: "1HGCM82633A123456", Year: 2025, Odometer: 300_000},
				{VIN: "2HGCM82633A654321", Year: 2021, Odometer: 90_000},
			},
		}
		body, _ := json.Marshal(vb)

		req := httptest.NewRequest(http.MethodPut, "/vehicles/bulk", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

		var got domain.VehiclesBulk
		err := json.NewDecoder(rec.Body).Decode(&got)
		assert.NoError(t, err)
		assert.Equal(t, vb.Vehicles[0].Year, got.Vehicles[0].Year)
		assert.Equal(t, vb.Vehicles[1].Odometer, got.Vehicles[1].Odometer)
	})

	// Invalid case
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/vehicles/bulk", bytes.NewReader([]byte(`{"vehicles":123}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
