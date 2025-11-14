package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// healthHandler provides a simple health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

// NewRouter sets up the HTTP routes for inspection operations
func NewRouter(handler *InspectionHandler) http.Handler {

	// Create a new router
	router := mux.NewRouter()

	// Health check endpoint
	router.HandleFunc("/health", healthHandler)

	// Inspection endpoints
	router.HandleFunc(
		"/inspections/inspect",
		func(w http.ResponseWriter, r *http.Request) {
			handler.InspectVehicle(w, r)
		},
	).Methods("POST")

	router.HandleFunc(
		"/inspections/get-build-data/{vin}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetBuildData(w, r)
		},
	).Methods("GET")

	return LoggingMiddleware(router)
}
