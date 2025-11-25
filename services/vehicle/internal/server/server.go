package server

import (
	"log/slog"
	"net/http"
	"time"

	vehiclehttp "github.com/alechekz/online-car-auction/services/vehicle/delivery/http"
	"github.com/alechekz/online-car-auction/services/vehicle/infrastructure"
	"github.com/alechekz/online-car-auction/services/vehicle/internal/logger"
	"github.com/alechekz/online-car-auction/services/vehicle/repository"
	"github.com/alechekz/online-car-auction/services/vehicle/usecase"
)

// Server represents the HTTP server for the Vehicle Service
type Server struct {
	httpServer *http.Server
}

// NewServer creates and configures a new Server instance with PostgreSQL repository
func NewServer(cfg *config) (*Server, error) {

	// dependencies
	var repo repository.VehicleRepository
	switch cfg.Repo {
	case "postgres":
		logger.Log.Info("using postgres vehicle repository")
		var err error
		repo, err = infrastructure.NewPostgresVehicleRepo(cfg.DatabaseURL)
		if err != nil {
			logger.Log.Error("failed to connect to postgres", slog.String("error", err.Error()))
		}
	default:
		logger.Log.Info("using in-memory vehicle repository")
		repo = infrastructure.NewMemoryVehicleRepo()
	}
	inspectionProvider, err := infrastructure.NewInspectionGRPCClient(cfg.InspectionURL)
	if err != nil {
		logger.Log.Error("failed to create inspection gRPC client", slog.String("error", err.Error()))
	}
	pricingProvider, err := infrastructure.NewPricingGRPCClient(cfg.PricingURL)
	if err != nil {
		logger.Log.Error("failed to create pricing gRPC client", slog.String("error", err.Error()))
	}
	logger.Log.Info("connected to dependencies", slog.String("inspection_url", cfg.InspectionURL), slog.String("pricing_url", cfg.PricingURL))
	uc := usecase.NewVehicleUC(repo, inspectionProvider, pricingProvider)
	handler := &vehiclehttp.VehicleHandler{UC: uc}
	mux := vehiclehttp.NewRouter(handler)

	return &Server{
		httpServer: &http.Server{
			Addr:              cfg.Address,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		},
	}, nil
}

// Start runs the HTTP server
func (s *Server) Start() error {
	logger.Log.Info("starting server", slog.String("addr", s.httpServer.Addr))
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server
func (s *Server) Stop() error {
	logger.Log.Info("shutting down server")
	return s.httpServer.Close()
}

// Handler returns the HTTP handler of the server
func (s *Server) Handler() http.Handler {
	return s.httpServer.Handler
}
