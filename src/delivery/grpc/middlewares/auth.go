package grpcm

import (
	"context"
	"strings"

	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (mw *Middleware) VerifyJWT() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if isPublicMethod(ctx) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization token")
		}

		token := strings.TrimPrefix(tokens[0], "Bearer ")
		if token == "" {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization token")
		}

		claims, ok := mw.authUsecase.VerifyToken(token, mw.cfg.SecretKey, authdom.Access)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		ctx = context.WithValue(ctx, cdom.Source, claims.Subject)
		return handler(ctx, req)
	}
}

func isPublicMethod(ctx context.Context) bool {

	publicMethods := []string{
		"/gobp.v1.GobpService/CreateToken",
		"/gobp.v1.GobpService/Ping",
		"/gobp.v1.GobpService/HealthCheck",
	}

	method, _ := grpc.Method(ctx)
	for _, publicMethod := range publicMethods {
		if method == publicMethod {
			return true
		}
	}
	return false
}
