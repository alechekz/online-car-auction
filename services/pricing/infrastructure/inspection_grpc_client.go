package infrastructure

import (
	"context"
	"time"

	pb "github.com/alechekz/online-car-auction/services/inspection/delivery/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InspectionGRPCClient is a gRPC client for the Inspection Service
type InspectionGRPCClient struct {
	client pb.InspectionServiceClient
	conn   *grpc.ClientConn
}

// NewInspectionGRPCClient creates a new InspectionGRPCClient instance
func NewInspectionGRPCClient(address string) (*InspectionGRPCClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewInspectionServiceClient(conn)
	return &InspectionGRPCClient{client: client, conn: conn}, nil
}

// Close closes the gRPC connection
func (c *InspectionGRPCClient) Close() error {
	return c.conn.Close()
}

// GetMsrp retrieves MSRP from the Inspection Service
func (c *InspectionGRPCClient) GetMsrp(vin string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.GetBuildDataRequest{Vin: vin}
	resp, err := c.client.GetBuildData(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Msrp, nil
}
