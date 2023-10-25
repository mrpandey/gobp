package migration

import (
	"os/exec"
	"strings"

	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/database"
	"github.com/mrpandey/gobp/src/util/projectpath"
)

func ApplyMigrations(cfg *util.Config, logger *util.StandardLogger) {
	dbUri := database.GetPgDbUri(cfg) + getDbUriParams(cfg)

	// create migration command
	cmd := exec.Command(`atlas`, `migrate`, `apply`, `--url`, dbUri, `--dir`, migrationDir)
	cmd.Dir = projectpath.Root // migration directory is at project root

	logger.Infof("Executing command: %s", cmd.String())

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Panicf("Failed to apply migrations. %v\n%s", err, out)
	}
}

func CheckMigrationStatus(cfg *util.Config, logger *util.StandardLogger) {
	dbUri := database.GetPgDbUri(cfg) + getDbUriParams(cfg)
	// create command
	cmd := exec.Command(
		`atlas`,
		`migrate`,
		`status`,
		`--url`,
		dbUri,
		`--dir`,
		migrationDir,
	)
	cmd.Dir = projectpath.Root

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Panicf("Failed to check migration status. %v\n%s", err, out)
	}

	ok := strings.HasPrefix(string(out), "Migration Status: OK")

	if !ok {
		logger.Panic("Migration status is not OK.")
	} else {
		logger.Info("Migration status OK.")

	}
}
