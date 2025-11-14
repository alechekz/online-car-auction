package server

import (
	"log/slog"
	"net/http"
	"time"

	vehiclehttp "github.com/alechekz/online-car-auction/services/inspection/delivery/http"
	"github.com/alechekz/online-car-auction/services/inspection/infrastructure"
	"github.com/alechekz/online-car-auction/services/inspection/internal/logger"
	"github.com/alechekz/online-car-auction/services/inspection/repository"
	"github.com/alechekz/online-car-auction/services/inspection/usecase"
)

// Server represents the HTTP server for the Inspection Service
type Server struct {
	httpServer *http.Server
}

// NewServer creates and configures a new Server instance with PostgreSQL repository
func NewServer(cfg *config) (*Server, error) {

	// dependencies
	var repo repository.InspectionRepository
	switch cfg.Repo {
	case "postgres":
		logger.Log.Info("using postgres inspection repository")
		var err error
		repo, err = infrastructure.NewPostgresInspectionRepo(cfg.DatabaseURL)
		if err != nil {
			logger.Log.Error("failed to connect to postgres", slog.String("error", err.Error()))
		}
	default:
		logger.Log.Info("using in-memory inspection repository")
		repo = infrastructure.NewMemoryInspectionRepo()
	}
	provider := infrastructure.NewNHTSABuildDataClient()
	uc := usecase.NewInspectionUC(repo, provider)
	handler := &vehiclehttp.InspectionHandler{UC: uc}
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
