package testuc

import (
	"testing"

	"github.com/mrpandey/gobp/src/internal/core/usecase"
	ucmock "github.com/mrpandey/gobp/src/internal/core/usecase/testutil/mocks"
)

func NewTestUseCase(t *testing.T) *usecase.UseCases {
	return &usecase.UseCases{
		Furniture: ucmock.NewFurnitureUseCaseInterface(t),
	}
}
