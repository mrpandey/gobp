package grpch

import (
	"context"

	pb_v1 "github.com/mrpandey/gobp/src/delivery/grpc/proto/gen/go/gobp/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GRPCServer) CreateToken(
	ctx context.Context,
	req *pb_v1.CreateTokenRequest,
) (*pb_v1.CreateTokenResponse, error) {
	response, err := s.useCases.Auth.CreateToken(ctx, req.ClientSlug, req.ClientSecret)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to generate token: %v", err)
	}

	pbResponse := &pb_v1.CreateTokenResponse{
		AccessToken: response.AccessToken,
		ExpiresIn:   uint64(response.AccessExpiresIn),
		TokenType:   response.TokenType,
	}

	return pbResponse, nil
}
