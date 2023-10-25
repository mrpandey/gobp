package main

import (
	"context"

	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/migration"
)

func main() {
	globalContext := context.Background()
	cfg := util.NewConfig()
	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)
	migration.GenerateMigrations(globalContext, cfg, logger)
}
