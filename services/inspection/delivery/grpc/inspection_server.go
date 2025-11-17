package grpc

import (
	"context"
	pb "github.com/alechekz/online-car-auction/services/inspection/delivery/grpc/proto"
	"github.com/alechekz/online-car-auction/services/inspection/domain"
	"github.com/alechekz/online-car-auction/services/inspection/usecase"
)

// InspectionServer implements the gRPC server for inspection service
type InspectionServer struct {
	pb.UnimplementedInspectionServiceServer
	uc usecase.InspectionUsecase
}

// NewInspectionServer creates a new InspectionServer instance
func NewInspectionServer(uc usecase.InspectionUsecase) *InspectionServer {
	return &InspectionServer{uc: uc}
}

// GetBuildData retrieves the build data for a vehicle by its VIN
func (s *InspectionServer) GetBuildData(ctx context.Context, req *pb.GetBuildDataRequest) (*pb.BuildDataResponse, error) {
	data, err := s.uc.GetBuildData(req.Vin)
	if err != nil {
		return nil, err
	}
	return &pb.BuildDataResponse{
		Vin:          data.VIN,
		Brand:        data.Brand,
		Engine:       data.Engine,
		Transmission: data.Transmission,
	}, nil
}

// InspectVehicle inspects a vehicle and returns its grade
func (s *InspectionServer) InspectVehicle(ctx context.Context, req *pb.InspectVehicleRequest) (*pb.InspectVehicleResponse, error) {
	v := &domain.Vehicle{
		VIN:             req.Vin,
		Odometer:        int(req.Odometer),
		Year:            int(req.Year),
		StrongScratches: req.StrongScratches,
		SmallScratches:  req.SmallScratches,
		ElectricFail:    req.ElectricFail,
		SuspensionFail:  req.SuspensionFail,
	}
	if err := s.uc.InspectVehicle(v); err != nil {
		return nil, err
	}
	return &pb.InspectVehicleResponse{
		Vin:   req.Vin,
		Grade: int32(v.Grade), //nolint:gosec // always between 1 and 50
	}, nil
}
