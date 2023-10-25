package grpcm

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Recovers from panic.
func (mw *Middleware) PanicRecovery() grpc.UnaryServerInterceptor {
	recoveryOption := grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		mw.logger.Error("panic: ", p)
		s := status.New(codes.Internal, "internal server error")
		return s.Err()
	})
	return grpc_recovery.UnaryServerInterceptor(recoveryOption)
}
