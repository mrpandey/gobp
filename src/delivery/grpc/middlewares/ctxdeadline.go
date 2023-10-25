package grpcm

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func (mw *Middleware) SetContextDeadline() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (_ interface{}, err error) {
		// Set context timeout if not set
		var newDeadline time.Time
		deadline, ok := ctx.Deadline()
		if !ok {
			// Here we use our own context timeout
			newDeadline = time.Now().Add(mw.cfg.ContextTimeout)
		} else {
			// Here we use propagated context timeout with subtraction of average network latency
			newDeadline = deadline.Add(mw.cfg.ContextTimeoutBuffer)
		}

		ctx, cancel := context.WithDeadline(ctx, newDeadline)
		// Release resources after request processed
		defer cancel()

		return handler(ctx, req)
	}
}
