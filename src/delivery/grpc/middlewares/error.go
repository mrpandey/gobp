package grpcm

import (
	"context"
	"encoding/json"
	"errors"

	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Convert errors to gRPC error.
func (mw *Middleware) ErrorHandler() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (_ interface{}, err error) {
		res, err := handler(ctx, req)
		if err == nil {
			return res, err
		}

		var jsonBytes []byte
		var jsonErr error
		var grpcStatusCode codes.Code
		var gobpErr gobperror.GobpError

		if errors.As(err, &gobpErr) {
			grpcStatusCode = gobpErr.GrpcCode
			jsonBytes, jsonErr = json.Marshal(gobpErr)
		} else {
			errStatus, ok := status.FromError(err)
			if ok {
				// It is already a gRPC error.
				grpcStatusCode = errStatus.Code()
				jsonBytes = []byte(errStatus.Message())
			} else {
				gobpErr = gobperror.ErrUnexpected
				grpcStatusCode = gobpErr.GrpcCode
				jsonBytes, jsonErr = json.Marshal(gobpErr)
			}
		}

		if jsonErr != nil {
			mw.logger.Errorf("Failed to marshal error response %v as json. %v", err, jsonErr)
			return res, status.Error(grpcStatusCode, "Failed to marshal error message.")
		}

		return res, status.Error(grpcStatusCode, string(jsonBytes))
	}
}
