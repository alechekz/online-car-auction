package server

import (
	"log/slog"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcDelivery "github.com/alechekz/online-car-auction/services/pricing/delivery/grpc"
	pb "github.com/alechekz/online-car-auction/services/pricing/delivery/grpc/proto"
	httpDelivery "github.com/alechekz/online-car-auction/services/pricing/delivery/http"

	"github.com/alechekz/online-car-auction/services/pricing/infrastructure"
	"github.com/alechekz/online-car-auction/services/pricing/internal/logger"
	"github.com/alechekz/online-car-auction/services/pricing/usecase"
)

// Server represents both HTTP and gRPC servers for the Pricing Service
type Server struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	grpcLis    net.Listener
}

// NewServer creates and configures a new Server instance
func NewServer(cfg *config) (*Server, error) {

	// Dependencies
	provider, err := infrastructure.NewInspectionGRPCClient(cfg.InspectionURL)
	if err != nil {
		return nil, err
	}
	uc := usecase.NewPricingUC(provider)

	// HTTP handler
	handler := &httpDelivery.PricingHandler{UC: uc}
	mux := httpDelivery.NewRouter(handler)

	// gRPC handler
	grpcSrv := grpc.NewServer()
	pb.RegisterPricingServiceServer(grpcSrv, grpcDelivery.NewPricingServer(uc))
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
