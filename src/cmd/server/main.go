package main

import (
	"context"

	grpcd "github.com/mrpandey/gobp/src/delivery/grpc"
	restd "github.com/mrpandey/gobp/src/delivery/rest"

	usecase "github.com/mrpandey/gobp/src/internal/core/usecase"
	"github.com/mrpandey/gobp/src/internal/repo"

	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/database"
	"github.com/mrpandey/gobp/src/util/migration"
)

func main() {
	globalContext := context.Background()

	cfg := util.NewConfig()

	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)

	logger.Info("Starting up api server...")

	logger.Info("Checking migration status...")
	migration.CheckMigrationStatus(cfg, logger)

	logger.Debug("Initializing postgres...")
	gormLogLevel := database.GetDefaultGormLogLevel(cfg.DGN)
	db := database.NewGormPostgresClient(globalContext, cfg, logger, gormLogLevel)

	logger.Debug("Initializing repo...")
	repos := repo.InitRepo(cfg, logger, db)

	logger.Debug("Initializing usecase...")
	useCases := usecase.InitUseCases(cfg, logger, repos)

	validator := util.NewRequestBodyValidator(logger)

	logger.Debug("Initializing HTTP server...")
	go restd.InitHTTPDelivery(globalContext, cfg, logger, validator, useCases)

	logger.Debug("Initializing GRPC server...")
	grpcd.InitGRPCDelivery(globalContext, cfg, logger, validator, useCases)

}
