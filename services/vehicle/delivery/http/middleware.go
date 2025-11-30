package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/alechekz/online-car-auction/services/vehicle/internal/logger"
)

// LoggingMiddleware logs details about each HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		logger.Log.Info("HTTP request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Duration("duration", duration),
		)
	})
}

// WriteError writes an error response based on the error type
func WriteError(w http.ResponseWriter, err error) {

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
