package grpch

import (
	"context"
	"time"

	pb_v1 "github.com/mrpandey/gobp/src/delivery/grpc/proto/gen/go/gobp/v1"
)

func (s *GRPCServer) Ping(_ context.Context, _ *pb_v1.PingRequest) (*pb_v1.PingResponse, error) {
	pingMsg := s.useCases.Health.Ping()
	return &pb_v1.PingResponse{Message: pingMsg}, nil
}

func (s *GRPCServer) HealthCheck(
	ctx context.Context,
	_ *pb_v1.HealthCheckRequest,
) (*pb_v1.HealthCheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbHealth := s.useCases.Health.HealthCheck(ctx)
	response := &pb_v1.HealthCheckResponse{
		PgReachable: dbHealth.PostgresReachable,
	}
	return response, nil
}
