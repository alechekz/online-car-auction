package server

import (
	"log/slog"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcDelivery "github.com/alechekz/online-car-auction/services/inspection/delivery/grpc"
	pb "github.com/alechekz/online-car-auction/services/inspection/delivery/grpc/proto"
	httpDelivery "github.com/alechekz/online-car-auction/services/inspection/delivery/http"

	"github.com/alechekz/online-car-auction/services/inspection/infrastructure"
	"github.com/alechekz/online-car-auction/services/inspection/internal/logger"
	"github.com/alechekz/online-car-auction/services/inspection/usecase"
)

// Server represents both HTTP and gRPC servers for the Inspection Service
type Server struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	grpcLis    net.Listener
}

// NewServer creates and configures a new Server instance
func NewServer(cfg *config) (*Server, error) {
	// dependencies
	provider := infrastructure.NewNHTSABuildDataClient()
	msrp := infrastructure.NewMockMSRPClient()
	uc := usecase.NewInspectionUC(provider, msrp)

	// HTTP handler
	handler := &httpDelivery.InspectionHandler{UC: uc}
	mux := httpDelivery.NewRouter(handler)

	// gRPC handler
	grpcSrv := grpc.NewServer()
	pb.RegisterInspectionServiceServer(grpcSrv, grpcDelivery.NewInspectionServer(uc))
	reflection.Register(grpcSrv)
	lis, err := net.Listen("tcp", cfg.GrpcAddress)
	if err != nil {
		return nil, err
	}

	// create server
	return &Server{
		httpServer: &http.Server{
			Addr:              cfg.HttpAddress,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		},
		grpcServer: grpcSrv,
		grpcLis:    lis,
	}, nil
}

// Start runs both HTTP and gRPC servers
func (s *Server) Start() error {
	// HTTP
	go func() {
		logger.Log.Info("starting HTTP server", slog.String("addr", s.httpServer.Addr))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error("http server error", slog.String("err", err.Error()))
		}
	}()

	// gRPC
	go func() {
		logger.Log.Info("starting gRPC server", slog.String("addr", s.grpcLis.Addr().String()))
		if err := s.grpcServer.Serve(s.grpcLis); err != nil {
			logger.Log.Error("grpc server error", slog.String("err", err.Error()))
		}
	}()

	return nil
}

// Stop gracefully shuts down both servers
func (s *Server) Stop() error {
	logger.Log.Info("shutting down servers")
	if err := s.httpServer.Close(); err != nil {
		return err
	}
	s.grpcServer.GracefulStop()
	return nil
}

// Handler returns the HTTP handler of the server
func (s *Server) Handler() http.Handler {
	return s.httpServer.Handler
}
