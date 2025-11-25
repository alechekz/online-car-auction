package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	vehiclehttp "github.com/alechekz/online-car-auction/services/vehicle/delivery/http"
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/infrastructure"
	"github.com/alechekz/online-car-auction/services/vehicle/internal/logger"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
)

// TestMain sets up the testing environment
func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

// newTestRouter creates a test HTTP router with in-memory dependencies
func newTestRouter() http.Handler {
	repo := infrastructure.NewMemoryVehicleRepo()
	inspectionProvider := &infrastructure.MockInspectionProvider{
		Data: &domain.Vehicle{
			VIN:          "1HGCM82633A004352",
			Brand:        "Kia",
			Engine:       "1.8L",
			Transmission: "Automatic",
		},
	}
	pricingProvider := &infrastructure.MockPricingProvider{}
	uc := usecase.NewVehicleUC(repo, inspectionProvider, pricingProvider)
	handler := &vehiclehttp.VehicleHandler{UC: uc}
	return vehiclehttp.NewRouter(handler)
}

// TestVehicleHandler_CreateVehicle tests the CreateVehicle HTTP handler
func TestVehicleHandler_CreateVehicle(t *testing.T) {
	router := newTestRouter()

	// Valid case
	t.Run("valid request", func(t *testing.T) {
		v := domain.Vehicle{
			VIN:      "1HGCM82633A123456",
			Year:     2020,
			Odometer: 15000,
			MSRP:     25000,
		}
		body, _ := json.Marshal(v)

		req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	// Invalid case
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader([]byte(`{"vin":123}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

// TestVehicleHandler_GetVehicle tests the GetVehicle HTTP handler
func TestVehicleHandler_GetVehicle(t *testing.T) {

	// Prepare router with a vehicle
	router := newTestRouter()
	v := domain.Vehicle{
		VIN:      "1HGCM82633A123456",
		Year:     2020,
		Odometer: 15000,
		MSRP:     25000,
	}
	body, _ := json.Marshal(v)
	req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Valid case
	t.Run("existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/vehicles/"+v.VIN, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var got domain.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, v.VIN, got.VIN)
	})

	// Invalid case
	t.Run("non-existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/vehicles/NONEXISTENTVIN", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

// TestVehicleHandler_UpdateVehicle tests the UpdateVehicle HTTP handler
func TestVehicleHandler_UpdateVehicle(t *testing.T) {

	// Prepare router with a vehicle
	router := newTestRouter()
	v := domain.Vehicle{
		VIN:      "1HGCM82633A123456",
		Year:     2020,
		Odometer: 15000,
		MSRP:     25000,
	}
	body, _ := json.Marshal(v)
	req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	updated := domain.Vehicle{
		Year:     2021,
		Odometer: 20000,
		MSRP:     27000,
	}
	body, _ = json.Marshal(updated)

	// Valid case
	t.Run("update existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/vehicles/"+v.VIN, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var got domain.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, int32(2021), got.Year)
		assert.Equal(t, int32(20000), got.Odometer)
	})

	// Invalid case
	t.Run("update non-existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/vehicles/NONEXISTENTVIN123", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

// TestVehicleHandler_DeleteVehicle tests the DeleteVehicle HTTP handler
func TestVehicleHandler_DeleteVehicle(t *testing.T) {

	// Prepare router with a vehicle
	router := newTestRouter()
	v := domain.Vehicle{
		VIN:      "1HGCM82633A123456",
		Year:     2020,
		Odometer: 15000,
		MSRP:     25000,
	}
	body, _ := json.Marshal(v)
	req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Valid case
	t.Run("delete existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/vehicles/"+v.VIN, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	// Invalid case
	t.Run("delete non-existing vehicle", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/vehicles/NONEXISTENTVIN", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

// TestVehicleHandler_ListVehicles tests the ListVehicles HTTP handler
func TestVehicleHandler_ListVehicles(t *testing.T) {

	// Prepare router with vehicles
	router := newTestRouter()
	v1 := domain.Vehicle{
		VIN:      "VIN12345678901234",
		Year:     2020,
		Odometer: 10000,
		MSRP:     20000,
	}
	v2 := domain.Vehicle{
		VIN:      "VIN23456789012345",
		Year:     2021,
		Odometer: 5000,
		MSRP:     22000,
	}
	for _, v := range []domain.Vehicle{v1, v2} {
		body, _ := json.Marshal(v)
		req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

	// Not empty list case
	t.Run("list all vehicles", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var got []domain.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Len(t, got, 2)

		vins := []string{got[0].VIN, got[1].VIN}
		assert.Contains(t, vins, v1.VIN)
		assert.Contains(t, vins, v2.VIN)
	})

	// Empty list case
	t.Run("list empty vehicles", func(t *testing.T) {
		router := newTestRouter()
		req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var got []domain.Vehicle
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Len(t, got, 0)
	})
}
