package fuc

import (
	"context"

	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	"github.com/mrpandey/gobp/src/util"
)

// Implement FurnitureUseCaseInterface.
type FurnitureUseCase struct {
	logger *util.StandardLogger
	txnMgr cdom.TxnManagerInterface
	fRepo  fdom.FurnitureRepoInterface
}

func NewFurnitureUseCase(
	logger *util.StandardLogger,
	txnMgr cdom.TxnManagerInterface,
	fRepo fdom.FurnitureRepoInterface,
) *FurnitureUseCase {
	return &FurnitureUseCase{
		logger: logger,
		txnMgr: txnMgr,
		fRepo:  fRepo,
	}
}

func (fuc *FurnitureUseCase) AddFurniture(
	ctx context.Context,
	req *fdom.AddFurnitureRequest,
) (*fdom.FurnitureID, error) {
	// Transaction not needed for single db operation. But just for demonstration.
	// Begin transaction.
	tx, rollback := fuc.txnMgr.Begin()
	defer rollback()

	furnitureID, err := fuc.fRepo.WithTx(tx).Create(ctx, req.Name, req.Type)
	if err != nil {
		return nil, err
	}

	// Commit transaction.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &fdom.FurnitureID{ID: furnitureID}, err
}

func (fuc *FurnitureUseCase) GetFurniture(ctx context.Context, req *fdom.FurnitureID) (*fdom.FurnitureRecord, error) {
	return fuc.fRepo.GetByID(ctx, req.ID)
}
