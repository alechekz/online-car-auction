package http

import (
	"encoding/json"
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
	"net/http"
	"strings"
)

// VehicleHandler handles HTTP requests for vehicle operations
type VehicleHandler struct {
	UC usecase.VehicleUsecase
}

// POST /vehicles
func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var v domain.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		WriteError(w, domain.ErrValidation)
		return
	}
	if err := h.UC.Create(&v); err != nil {
		WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(v)
}

// GET /vehicles/{vin}
func (h *VehicleHandler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/vehicles/")
	v, err := h.UC.Get(vin)
	if err != nil {
		WriteError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(v)
}

// PUT /vehicles/{vin}
func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/vehicles/")
	var v domain.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		WriteError(w, domain.ErrValidation)
		return
	}
	v.VIN = vin
	if err := h.UC.Update(&v); err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

// DELETE /vehicles/{vin}
func (h *VehicleHandler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	vin := strings.TrimPrefix(r.URL.Path, "/vehicles/")
	if err := h.UC.Delete(vin); err != nil {
		WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "deleted", "vin": vin})
}

// GET /vehicles
func (h *VehicleHandler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.UC.List()
	if err != nil {
		WriteError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(vehicles)
}
