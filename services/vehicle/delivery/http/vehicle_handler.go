package http

import (
	"encoding/json"
	"errors"
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
	"net/http"
	"strings"
)

// VehicleHandler handles HTTP requests for vehicle operations
type VehicleHandler struct {
	UC usecase.VehicleUsecase
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

// POST /vehicles
func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var v domain.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		writeError(w, domain.ErrValidation)
		return
	}
	if err := h.UC.CreateVehicle(&v); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(v)
}

// GET /vehicles/{vin}
func (h *VehicleHandler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/vehicles/")
	v, err := h.UC.GetVehicle(vin)
	if err != nil {
		writeError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(v)
}

// PUT /vehicles/{vin}
func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/vehicles/")
	var v domain.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		writeError(w, domain.ErrValidation)
		return
	}
	v.VIN = vin
	if err := h.UC.UpdateVehicle(&v); err != nil {
		writeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

// DELETE /vehicles/{vin}
func (h *VehicleHandler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/vehicles/")
	if err := h.UC.DeleteVehicle(vin); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "deleted", "vin": vin})
}

// GET /vehicles
func (h *VehicleHandler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.UC.ListVehicles()
	if err != nil {
		writeError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(vehicles)
}
