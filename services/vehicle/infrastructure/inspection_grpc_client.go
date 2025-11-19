package infrastructure

import (
	"context"
	"time"

	pb "github.com/alechekz/online-car-auction/services/inspection/delivery/grpc/proto"
	"github.com/alechekz/online-car-auction/services/vehicle/domain"
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

// InspectVehicle sends an inspection request to the Inspection Service
func (c *InspectionGRPCClient) InspectVehicle(v *domain.Vehicle) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.InspectVehicleRequest{
		Vin:             v.VIN,
		Year:            v.Year,
		Odometer:        v.Odometer,
		SmallScratches:  v.SmallScratches,
		StrongScratches: v.StrongScratches,
		ElectricFail:    v.ElectricFail,
		SuspensionFail:  v.SuspensionFail,
	}

	resp, err := c.client.InspectVehicle(ctx, req)
	if err != nil {
		return err
	}
	v.Grade = int(resp.Grade)
	return nil
}

// GetBuildData retrieves the build data of a vehicle from the Inspection Service
func (c *InspectionGRPCClient) GetBuildData(vin string) (*domain.Vehicle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.GetBuildDataRequest{Vin: vin}
	resp, err := c.client.GetBuildData(ctx, req)
	if err != nil {
		return nil, err
	}
	return &domain.Vehicle{
		VIN:          resp.Vin,
		Brand:        resp.Brand,
		Engine:       resp.Engine,
		Transmission: resp.Transmission,
		MSRP:         resp.Msrp,
	}, nil
}
