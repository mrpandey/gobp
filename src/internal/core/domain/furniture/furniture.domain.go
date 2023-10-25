package fdom

import (
	"context"

	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"

	"github.com/go-playground/validator/v10"
)

type FurnitureType string

const (
	Table    FurnitureType = "table"
	Chair    FurnitureType = "chair"
	Wardrobe FurnitureType = "wardrobe"
)

type FurnitureUseCaseInterface interface {
	AddFurniture(ctx context.Context, req *AddFurnitureRequest) (*FurnitureID, error)
	GetFurniture(ctx context.Context, req *FurnitureID) (*FurnitureRecord, error)
}

type FurnitureRepoInterface interface {
	WithTx(tm cdom.TxnManagerInterface) FurnitureRepoInterface
	Create(ctx context.Context, name string, typ FurnitureType) (furnitureID uint, err error)
	GetByID(ctx context.Context, id uint) (*FurnitureRecord, error)
}

func FurnitureTypeValidator(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(FurnitureType)
	switch value {
	case Table, Chair, Wardrobe:
		return true
	default:
		return false
	}
}
