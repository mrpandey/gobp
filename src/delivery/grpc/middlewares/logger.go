package grpcm

import (
	"context"
	"time"

	"github.com/mrpandey/gobp/src/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (mw *Middleware) LogRequest(logger *util.StandardLogger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (_ interface{}, err error) {

		startTime := time.Now()

		resposne, err := handler(ctx, req)

		duration := time.Since(startTime).Milliseconds()

		errStatus, _ := status.FromError(err)

		logger.Infof("Unary grpc call %v returned %v in %v ms", info.FullMethod, errStatus.Code(), duration)

		return resposne, err
	}
}
