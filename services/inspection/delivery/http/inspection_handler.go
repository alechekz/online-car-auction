package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/alechekz/online-car-auction/services/inspection/domain"
	"github.com/alechekz/online-car-auction/services/inspection/usecase"
)

// InspectionHandler handles HTTP requests for inspection operations
type InspectionHandler struct {
	UC usecase.InspectionUsecase
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

// POST /inspections/inspect
func (h *InspectionHandler) InspectVehicle(w http.ResponseWriter, r *http.Request) {
	var v domain.Inspection
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		writeError(w, domain.ErrValidation)
		return
	}
	if err := h.UC.InspectVehicle(&v); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(v)
}

// GET /inspections/get-build-data/{vin}
func (h *InspectionHandler) GetBuildData(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/inspections/get-build-data/")
	buildData, err := h.UC.GetBuildData(vin)
	if err != nil {
		writeError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(buildData)
}
