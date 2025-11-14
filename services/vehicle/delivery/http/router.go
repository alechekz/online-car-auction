package http

import "net/http"

// healthHandler provides a simple health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

// NewRouter sets up the HTTP routes for vehicle operations
func NewRouter(handler *VehicleHandler) http.Handler {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", healthHandler)

	// /vehicles (POST, GET)
	mux.HandleFunc("/vehicles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateVehicle(w, r)
		case http.MethodGet:
			handler.ListVehicles(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /vehicles/{vin} (GET, PUT, DELETE)
	mux.HandleFunc("/vehicles/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetVehicle(w, r)
		case http.MethodPut:
			handler.UpdateVehicle(w, r)
		case http.MethodDelete:
			handler.DeleteVehicle(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return LoggingMiddleware(mux)
}
