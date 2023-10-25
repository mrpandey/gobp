package grpcd

import (
	"context"
	"fmt"
	"net"

	grpch "github.com/mrpandey/gobp/src/delivery/grpc/handlers"
	grpcm "github.com/mrpandey/gobp/src/delivery/grpc/middlewares"
	pb_v1 "github.com/mrpandey/gobp/src/delivery/grpc/proto/gen/go/gobp/v1"

	"github.com/mrpandey/gobp/src/internal/core/usecase"
	"github.com/mrpandey/gobp/src/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitGRPCDelivery(
	_ context.Context,
	cfg *util.Config,
	logger *util.StandardLogger,
	validator *util.Validator,
	useCases *usecase.UseCases,
) {
	grpcListener, err := getGRPCListener(cfg)
	if err != nil {
		logger.Panic(err)
	}

	disableCallerLog := true
	requestLogger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)
	middleware := grpcm.NewMiddleware(cfg, logger, useCases.Auth)

	baseServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.PanicRecovery(),
		middleware.LogRequest(requestLogger),
		middleware.VerifyJWT(),
		middleware.SetContextDeadline(),
		middleware.ErrorHandler(),
	))

	gobpServer := grpch.NewGRPCServer(logger, validator, useCases)
	pb_v1.RegisterGobpServiceServer(baseServer, gobpServer)
	reflection.Register(baseServer)

	logger.Infof("gRPC server running on %s", grpcListener.Addr())
	err = baseServer.Serve(grpcListener)
	if err != nil {
		logger.Panic(err)
	}
}

func getGRPCListener(cfg *util.Config) (net.Listener, error) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort)
	grpcListener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return grpcListener, nil
}
