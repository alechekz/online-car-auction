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

// NewRouter sets up the HTTP routes for the pricing service
func NewRouter(handler *PricingHandler) http.Handler {

	// Create a new router
	router := mux.NewRouter()

	// Health check endpoint
	router.HandleFunc("/health", healthHandler)

	// Pricing endpoints
	router.HandleFunc(
		"/pricing/get-recommended-price",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetRecommendedPrice(w, r)
		},
	).Methods("POST")

	return LoggingMiddleware(router)
}
