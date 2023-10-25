package grpch

import (
	pb_v1 "github.com/mrpandey/gobp/src/delivery/grpc/proto/gen/go/gobp/v1"
	"github.com/mrpandey/gobp/src/internal/core/usecase"
	"github.com/mrpandey/gobp/src/util"
)

type GRPCServer struct {
	pb_v1.UnimplementedGobpServiceServer
	logger    *util.StandardLogger
	validator *util.Validator
	useCases  *usecase.UseCases
}

// Implements pb_v1.GobpServiceServer interface.
func NewGRPCServer(
	logger *util.StandardLogger,
	validator *util.Validator,
	useCases *usecase.UseCases,
) (server *GRPCServer) {
	server = &GRPCServer{
		logger:    logger,
		validator: validator,
		useCases:  useCases,
	}
	return server
}
