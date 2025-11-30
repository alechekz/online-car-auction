package http

import (
	"encoding/json"
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
	"net/http"
)

// VehiclesBulkHandler handles HTTP requests for bulk vehicle operations
type VehiclesBulkHandler struct {
	UC usecase.VehiclesBulkUsecase
}

// POST /vehicles/bulk
func (h *VehiclesBulkHandler) CreateVehiclesBulk(w http.ResponseWriter, r *http.Request) {
	var vb domain.VehiclesBulk
	if err := json.NewDecoder(r.Body).Decode(&vb); err != nil {
		WriteError(w, domain.ErrValidation)
		return
	}
	if err := h.UC.Create(&vb); err != nil {
		WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(vb)
}

// PUT /vehicles/bulk
func (h *VehiclesBulkHandler) UpdateVehiclesBulk(w http.ResponseWriter, r *http.Request) {
	var vb domain.VehiclesBulk
	if err := json.NewDecoder(r.Body).Decode(&vb); err != nil {
		WriteError(w, domain.ErrValidation)
		return
	}
	if err := h.UC.Update(&vb); err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(vb)
}
