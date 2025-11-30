package http

import "net/http"

// healthHandler provides a simple health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

// regVehiclesRoutes registers vehicle routes
func regVehiclesRoutes(mux *http.ServeMux, handler *VehicleHandler) {

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
}

// regBulkRoutes registers bulk vehicle routes
func regBulkRoutes(mux *http.ServeMux, handler *VehiclesBulkHandler) {

	// vehicles/bulk (POST, PUT)
	mux.HandleFunc("/vehicles/bulk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateVehiclesBulk(w, r)
		case http.MethodPut:
			handler.UpdateVehiclesBulk(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// NewRouter sets up the HTTP routes for vehicle operations
func NewRouter(handler *VehicleHandler, bulkHandler *VehiclesBulkHandler) http.Handler {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", healthHandler)

	// Register vehicle routes
	regVehiclesRoutes(mux, handler)

	// Register bulk vehicle routes
	regBulkRoutes(mux, bulkHandler)

	// Wrap with logging middleware and exit
	return LoggingMiddleware(mux)
}
