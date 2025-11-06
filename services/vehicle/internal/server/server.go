package server

import (
	"fmt"
	"net/http"

	vehiclehttp "github.com/alechekz/online-car-auction/services/vehicle/delivery/http"
	"github.com/alechekz/online-car-auction/services/vehicle/infrastructure"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
)

// Server represents the HTTP server for the Vehicle Service
type Server struct {
	httpServer *http.Server
}

// NewServer creates and configures a new Server instance
func NewServer(cfg *config) *Server {

	// dependencies
	repo := infrastructure.NewMemoryVehicleRepo()
	uc := usecase.NewVehicleUC(repo)
	handler := &vehiclehttp.VehicleHandler{UC: uc}
	mux := vehiclehttp.NewRouter(handler)

	// create http server
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Address,
			Handler: mux,
		},
	}
}

// Start runs the HTTP server
func (s *Server) Start() error {
	fmt.Printf("Vehicle Service started on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server
func (s *Server) Stop() error {
	fmt.Println("Shutting down server...")
	return s.httpServer.Close()
}

// Handler returns the HTTP handler of the server
func (s *Server) Handler() http.Handler {
	return s.httpServer.Handler
}
