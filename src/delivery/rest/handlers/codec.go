package resth

import (
	"net/http"
	"strconv"

	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	"github.com/mrpandey/gobp/src/util"

	"github.com/go-chi/chi"
)

func decodeGetFurnitureRequest(
	logger *util.StandardLogger,
	r *http.Request,
	v *util.Validator,
) (*fdom.FurnitureID, error) {
	furnitureID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, gobperror.NewInvalidUrlParamTypeError("id")
	}

	req := fdom.FurnitureID{ID: uint(furnitureID)}

	err = v.ValidateStruct(logger, req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func decodeAddFurnitureRequest(
	logger *util.StandardLogger,
	r *http.Request,
	v *util.Validator,
) (*fdom.AddFurnitureRequest, error) {
	req := fdom.AddFurnitureRequest{}
	err := decodeJSON(logger, r.Body, &req)
	if err != nil {
		return nil, err
	}

	err = v.ValidateStruct(logger, req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}
