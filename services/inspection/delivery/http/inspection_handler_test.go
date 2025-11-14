package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	inspectionhttp "github.com/alechekz/online-car-auction/services/inspection/delivery/http"
	"github.com/alechekz/online-car-auction/services/inspection/domain"
	"github.com/alechekz/online-car-auction/services/inspection/infrastructure"
	"github.com/alechekz/online-car-auction/services/inspection/internal/logger"
	"github.com/alechekz/online-car-auction/services/inspection/usecase"
)

// TestMain sets up the testing environment
func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

// newTestRouter creates a test HTTP router with in-memory dependencies
func newTestRouter() http.Handler {
	repo := infrastructure.NewMemoryInspectionRepo()
	provider := infrastructure.NewNHTSABuildDataClient()
	uc := usecase.NewInspectionUC(repo, provider)
	handler := &inspectionhttp.InspectionHandler{UC: uc}
	return inspectionhttp.NewRouter(handler)
}

// TestInspectionHandler_InspectVehicle tests the InspectVehicle HTTP handler
func TestHandler_InspectVehicle(t *testing.T) {
	router := newTestRouter()

	// Valid case
	t.Run("valid request", func(t *testing.T) {
		v := domain.Inspection{
			VIN:      "1HGCM82633A123456",
			Year:     2020,
			Odometer: 15000,
		}
		body, _ := json.Marshal(v)

		req := httptest.NewRequest(http.MethodPost, "/inspections/inspect", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	// Invalid case
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/inspections/inspect", bytes.NewReader([]byte(`{"vin":123}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

// TestInspectionHandler_GetBuildData tests the GetBuildData HTTP handler
func TestHandler_GetBuildData(t *testing.T) {
	router := newTestRouter()

	// Valid case
	t.Run("valid VIN", func(t *testing.T) {
		vin := "1HGCM82633A004352"
		req := httptest.NewRequest(http.MethodGet, "/inspections/get-build-data/"+vin, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	// Invalid case
	t.Run("invalid VIN", func(t *testing.T) {
		vin := "123"
		req := httptest.NewRequest(http.MethodGet, "/inspections/get-build-data/"+vin, nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
