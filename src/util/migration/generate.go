package migration

import (
	"context"
	"os/exec"

	authrepo "github.com/mrpandey/gobp/src/internal/repo/auth"
	frepo "github.com/mrpandey/gobp/src/internal/repo/furniture"

	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/database"
	"github.com/mrpandey/gobp/src/util/projectpath"

	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

const (
	migrationDir = "file://migrations?format=golang-migrate"
)

// Generates schema migration files.
func GenerateMigrations(ctx context.Context, cfg *util.Config, logger *util.StandardLogger) {
	cfg.PGDB = "postgres"
	defaultDB := database.NewGormPostgresClient(ctx, cfg, logger, glogger.Warn)
	dbUriParams := getDbUriParams(cfg)

	// will have the desired schema
	desiredSchemaDBName := "temp_migration_desired"
	database.DropPgDB(ctx, logger, defaultDB, desiredSchemaDBName)
	database.CreatePgDB(ctx, logger, defaultDB, desiredSchemaDBName)
	defer database.DropPgDB(ctx, logger, defaultDB, desiredSchemaDBName)

	// atlas dev db
	devDBName := "temp_migration_dev"
	database.DropPgDB(ctx, logger, defaultDB, devDBName)
	database.CreatePgDB(ctx, logger, defaultDB, devDBName)
	defer database.DropPgDB(ctx, logger, defaultDB, devDBName)

	// migration files (CURRENT STATE) will be applied on devDB
	cfg.PGDB = devDBName
	devDBUri := database.GetPgDbUri(cfg) + dbUriParams

	cfg.PGDB = desiredSchemaDBName
	desiredDBUri := database.GetPgDbUri(cfg) + dbUriParams

	desiredGormDB := database.NewGormPostgresClient(ctx, cfg, logger, glogger.Warn)
	desiredSqlDB := database.GetGormConnPool(logger, desiredGormDB)
	defer desiredSqlDB.Close()

	// create gorm models (DESIRED STATE) in desiredDB
	autoMigrate(ctx, logger, desiredGormDB)

	// create command to generate migration files by looking at diff between desired and current state
	cmd := exec.Command(
		`atlas`,
		`migrate`,
		`diff`,
		`auto_generated`,
		`--dir`,
		migrationDir,
		`--dev-url`,
		devDBUri,
		`--to`,
		desiredDBUri,
	)
	cmd.Dir = projectpath.Root

	logger.Infof("Running command: %s", cmd.String())

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Panicf("Failed to generate migrations. %v\n%s", err, out)
	}
}

func autoMigrate(ctx context.Context, logger *util.StandardLogger, db *gorm.DB) {
	err := db.WithContext(ctx).AutoMigrate(
		authrepo.ClientCred{},
		frepo.Furniture{},
	)
	if err != nil {
		logger.Panicf("Gorm automigration failed. %s", err)
	}
	logger.Info("Gorm AutoMigrate successful.")
}

func getDbUriParams(cfg *util.Config) string {
	dbURIParams := "?search_path=public"

	if cfg.PgSslMode == "disable" {
		dbURIParams += "&sslmode=disable"
	}

	return dbURIParams
}
