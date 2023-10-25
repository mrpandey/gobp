package healthrepo

import (
	"context"
	"database/sql"

	repoerror "github.com/mrpandey/gobp/src/internal/repo/error"
	"github.com/mrpandey/gobp/src/util"
)

// healthRepo implements HealthRepoInterface.
type healthRepo struct {
	config *util.Config
	logger *util.StandardLogger
	sqlDB  *sql.DB
}

func NewHealthRepo(config *util.Config, logger *util.StandardLogger, db *sql.DB) *healthRepo {
	return &healthRepo{
		config: config,
		logger: logger,
		sqlDB:  db,
	}
}

func (hr *healthRepo) PingPostgres(_ context.Context) error {
	if hr.sqlDB == nil {
		err := repoerror.ErrNoDbConn
		hr.logger.Errorf("%v", err)
		return err
	}
	err := hr.sqlDB.Ping()
	if err != nil {
		hr.logger.Errorf("%v", err)
		return err
	}
	return nil
}
