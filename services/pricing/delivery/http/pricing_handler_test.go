package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	pricinghttp "github.com/alechekz/online-car-auction/services/pricing/delivery/http"
	"github.com/alechekz/online-car-auction/services/pricing/domain"
	"github.com/alechekz/online-car-auction/services/pricing/infrastructure"
	"github.com/alechekz/online-car-auction/services/pricing/internal/logger"
	"github.com/alechekz/online-car-auction/services/pricing/usecase"
)

// TestMain sets up the testing environment
func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

// newTestRouter creates a test HTTP router with in-memory dependencies
func newTestRouter() http.Handler {
	provider := &infrastructure.MockInspectionProvider{
		Data: &domain.Vehicle{
			VIN:  "1HGCM82633A004352",
			Msrp: 99_000,
		},
	}
	uc := usecase.NewPricingUC(provider)
	handler := &pricinghttp.PricingHandler{UC: uc}
	return pricinghttp.NewRouter(handler)
}

// TestPricingHandler_GetRecommendedPrice tests the GetRecommendedPrice HTTP handler
func TestHandler_GetRecommendedPrice(t *testing.T) {
	router := newTestRouter()

	// Valid case
	t.Run("valid request", func(t *testing.T) {
		v := domain.Vehicle{
			VIN:      "1HGCM82633A123456",
			Grade:    47,
			Odometer: 30_000,
		}
		body, _ := json.Marshal(v)

		req := httptest.NewRequest(http.MethodPost, "/pricing/get-recommended-price", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	// Invalid case
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/pricing/get-recommended-price", bytes.NewReader([]byte(`{"vin":123}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
