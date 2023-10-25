package resth

import (
	"net/http"
)

func (s *RestServer) GetFurniture(w http.ResponseWriter, r *http.Request) {
	req, err := decodeGetFurnitureRequest(s.logger, r, s.validator)
	if err != nil {
		SendErrorResponse(s.logger, w, err)
		return
	}

	furniture, err := s.useCases.Furniture.GetFurniture(r.Context(), req)
	if err != nil {
		SendErrorResponse(s.logger, w, err)
		return
	}

	sendJSONResponse(s.logger, w, http.StatusOK, furniture)
}

func (s *RestServer) AddFurniture(w http.ResponseWriter, r *http.Request) {
	req, err := decodeAddFurnitureRequest(s.logger, r, s.validator)
	if err != nil {
		SendErrorResponse(s.logger, w, err)
		return
	}

	furniture, err := s.useCases.Furniture.AddFurniture(r.Context(), req)
	if err != nil {
		SendErrorResponse(s.logger, w, err)
		return
	}

	sendJSONResponse(s.logger, w, http.StatusCreated, furniture)
}
