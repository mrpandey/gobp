package testrepo

import (
	"context"
	"testing"

	"github.com/mrpandey/gobp/src/internal/repo"
	repomock "github.com/mrpandey/gobp/src/internal/repo/testutil/mocks"
	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/database"
	"github.com/mrpandey/gobp/src/util/migration"

	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func NewTestRepo(t *testing.T) *repo.Repo {
	return &repo.Repo{
		Health:     repomock.NewHealthRepoInterface(t),
		TxnManager: repomock.NewTxnManagerInterface(t),
		Furniture:  repomock.NewFurnitureRepoInterface(t),
	}
}

// Creates test-db with name "gobptest" and applies migrations.
// Returns client for test-db, config, and teardown function that must be deferred.
func SetupTestDB() (*gorm.DB, *util.Config, func()) {
	testDBName := "gobptest"

	cfg := util.NewConfig()
	cfg.PGDB = "postgres"
	cfg.LogLevel = "warn"
	disableCallerLog := false

	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)
	pgClient := database.NewGormPostgresClient(context.TODO(), cfg, logger, glogger.Warn)

	database.DropPgDB(context.TODO(), logger, pgClient, testDBName)   // drop database if exists
	database.CreatePgDB(context.TODO(), logger, pgClient, testDBName) // create test database

	// new config with different database
	testCfg := util.NewConfig()
	testCfg.PGDB = testDBName

	testDBClient := database.NewGormPostgresClient(context.TODO(), testCfg, logger, glogger.Warn)
	migration.ApplyMigrations(testCfg, logger) // create tables in test database

	return testDBClient, testCfg, func() {
		// teardown
		database.GetGormConnPool(logger, testDBClient).Close()
		database.DropPgDB(context.TODO(), logger, pgClient, testDBName)
	}
}
