package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alechekz/online-car-auction/services/pricing/domain"
	"github.com/alechekz/online-car-auction/services/pricing/usecase"
)

// PricingHandler handles HTTP requests related to vehicle pricing
type PricingHandler struct {
	UC usecase.PricingUsecase
}

// writeError writes an error response based on the error type
func writeError(w http.ResponseWriter, err error) {

	// Default to internal server error
	status := http.StatusInternalServerError
	msg := "internal server error"

	// Determine specific error type
	switch {
	case errors.Is(err, domain.ErrValidation):
		status = http.StatusBadRequest
		msg = err.Error()
	case errors.Is(err, domain.ErrNotFound):
		status = http.StatusNotFound
		msg = err.Error()
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// POST /pricing/get-recommended-price
func (h *PricingHandler) GetRecommendedPrice(w http.ResponseWriter, r *http.Request) {
	var v domain.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		writeError(w, domain.ErrValidation)
		return
	}
	err := h.UC.GetRecommendedPrice(&v)
	if err != nil {
		writeError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]uint64{"recommended_price": v.Price})
}
