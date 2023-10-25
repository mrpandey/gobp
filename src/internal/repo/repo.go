package repo

import (
	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	healthdom "github.com/mrpandey/gobp/src/internal/core/domain/health"

	authrepo "github.com/mrpandey/gobp/src/internal/repo/auth"
	frepo "github.com/mrpandey/gobp/src/internal/repo/furniture"
	healthrepo "github.com/mrpandey/gobp/src/internal/repo/health"
	txnmgr "github.com/mrpandey/gobp/src/internal/repo/txnmanager"

	"github.com/mrpandey/gobp/src/util"
	"github.com/mrpandey/gobp/src/util/database"

	"gorm.io/gorm"
)

type Repo struct {
	Auth       authdom.AuthRepoInterface
	Health     healthdom.HealthRepoInterface
	TxnManager cdom.TxnManagerInterface
	Furniture  fdom.FurnitureRepoInterface
}

func InitRepo(config *util.Config, logger *util.StandardLogger, db *gorm.DB) *Repo {
	sqlDB := database.GetGormConnPool(logger, db)
	return &Repo{
		Auth:       authrepo.NewAuthRepo(logger, db),
		Health:     healthrepo.NewHealthRepo(config, logger, sqlDB),
		TxnManager: txnmgr.NewGormTxnManager(logger, db),
		Furniture:  frepo.NewFurnitureRepo(logger, db),
	}
}
