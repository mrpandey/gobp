package frepo

import (
	"context"
	"errors"

	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	"github.com/mrpandey/gobp/src/util"

	"gorm.io/gorm"
)

// Implements FurnitureRepoInterface.
type FurnitureRepo struct {
	db     *gorm.DB
	logger *util.StandardLogger
}

func NewFurnitureRepo(
	logger *util.StandardLogger,
	db *gorm.DB,
) *FurnitureRepo {
	return &FurnitureRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *FurnitureRepo) WithTx(tm cdom.TxnManagerInterface) fdom.FurnitureRepoInterface {
	return &FurnitureRepo{
		db:     tm.GetTx(),
		logger: repo.logger,
	}
}

func (repo *FurnitureRepo) Create(
	ctx context.Context,
	name string,
	typ fdom.FurnitureType,
) (furnitureID uint, err error) {
	furniture := Furniture{
		Name: name,
		Type: typ,
	}

	err = repo.db.WithContext(ctx).Create(&furniture).Error
	if err != nil {
		repo.logger.Errorf("Failed to create new furniture entry in database. %v", err)
		return 0, err
	}

	return furniture.ID, nil
}

func (repo *FurnitureRepo) GetByID(ctx context.Context, id uint) (*fdom.FurnitureRecord, error) {
	furniture := Furniture{}

	// TODO: select specific fields only
	err := repo.db.WithContext(ctx).First(&furniture, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gobperror.ErrRecordNotFound
		}
		repo.logger.Errorf("Failed to get furniture with id = %v from database. %v", id, err)
		return nil, err
	}

	return &fdom.FurnitureRecord{
		ID:   furniture.ID,
		Name: furniture.Name,
		Type: furniture.Type,
	}, nil
}
