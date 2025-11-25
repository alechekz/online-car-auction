package grpc

import (
	"context"
	pb "github.com/alechekz/online-car-auction/services/pricing/delivery/grpc/proto"
	"github.com/alechekz/online-car-auction/services/pricing/domain"
	"github.com/alechekz/online-car-auction/services/pricing/usecase"
)

// PricingServer implements the gRPC server for pricing service
type PricingServer struct {
	pb.UnimplementedPricingServiceServer
	uc usecase.PricingUsecase
}

// NewPricingServer creates a new PricingServer instance
func NewPricingServer(uc usecase.PricingUsecase) *PricingServer {
	return &PricingServer{uc: uc}
}

// GetRecommendedPrice retrieves the recommended price for a vehicle by its VIN
func (s *PricingServer) GetRecommendedPrice(ctx context.Context, req *pb.PriceRequest) (*pb.PriceResponse, error) {
	v := &domain.Vehicle{
		VIN:           req.Vin,
		Odometer:      int(req.Odometer),
		Grade:         int(req.Grade),
		ExteriorColor: req.ExteriorColor,
		InteriorColor: req.InteriorColor,
	}
	err := s.uc.GetRecommendedPrice(v)
	if err != nil {
		return nil, err
	}
	return &pb.PriceResponse{
		Price: v.Price,
	}, nil
}
