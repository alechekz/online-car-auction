package infrastructure

import (
	"context"
	"time"

	pb "github.com/alechekz/online-car-auction/services/pricing/delivery/grpc/proto"
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PricingGRPCClient is a gRPC client for the Pricing Service
type PricingGRPCClient struct {
	client pb.PricingServiceClient
	conn   *grpc.ClientConn
}

// NewPricingGRPCClient creates a new PricingGRPCClient instance
func NewPricingGRPCClient(address string) (*PricingGRPCClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewPricingServiceClient(conn)
	return &PricingGRPCClient{client: client, conn: conn}, nil
}

// Close closes the gRPC connection
func (c *PricingGRPCClient) Close() error {
	return c.conn.Close()
}

// GetRecommendedPrice sends price request to the Pricing Service
func (c *PricingGRPCClient) GetRecommendedPrice(v *domain.Vehicle) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req := &pb.PriceRequest{
		Vin:           v.VIN,
		Odometer:      v.Odometer,
		Grade:         int32(v.Grade), // nolint:gosec
		ExteriorColor: v.ExteriorColor,
		InteriorColor: v.InteriorColor,
	}
	resp, err := c.client.GetRecommendedPrice(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Price, nil
}
