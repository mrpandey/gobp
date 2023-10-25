package txnmgr

import (
	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	"github.com/mrpandey/gobp/src/util"

	"gorm.io/gorm"
)

// Implements TxnManager interface.
type GormTxnManager struct {
	gormTx *gorm.DB
	logger *util.StandardLogger
}

func NewGormTxnManager(logger *util.StandardLogger, db *gorm.DB) *GormTxnManager {
	return &GormTxnManager{
		gormTx: db,
		logger: logger,
	}
}

func (tm *GormTxnManager) GetTx() *gorm.DB {
	return tm.gormTx
}

func (tm *GormTxnManager) Begin() (cdom.TxnManagerInterface, func()) {
	tx := tm.gormTx.Begin()
	return &GormTxnManager{
			gormTx: tx,
			logger: tm.logger,
		}, func() {
			tx.Rollback()
		}
}

func (tm *GormTxnManager) Commit() error {
	err := tm.gormTx.Commit().Error
	if err != nil {
		tm.logger.Errorf("Could not commit transaction in database. %v", err)
	}
	return err
}

func (tm *GormTxnManager) Rollback() error {
	err := tm.gormTx.Rollback().Error
	if err != nil {
		tm.logger.Errorf("Could not rollback transaction in database. %v", err)
	}
	return err
}
