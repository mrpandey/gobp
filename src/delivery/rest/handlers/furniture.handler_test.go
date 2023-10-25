package resth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	testrest "github.com/mrpandey/gobp/src/delivery/rest/testutil"
	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	ucmock "github.com/mrpandey/gobp/src/internal/core/usecase/testutil/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFurnitureRequestValidation(t *testing.T) {

	subtests := map[string]struct {
		typ            fdom.FurnitureType
		name           string
		expectedStatus int
		expectedError  gobperror.GobpError
	}{
		"CheckFurnitureType": {
			typ:            "Table",
			name:           "Table for Test",
			expectedStatus: http.StatusBadRequest,
			expectedError:  gobperror.NewValidationError("type: invalid furniture type"),
		},
		"TypeRequired": {
			name:           "Another Table",
			expectedStatus: http.StatusBadRequest,
			expectedError:  gobperror.NewValidationError("type is a required field"),
		},
		"NameRequired": {
			typ:            fdom.Table,
			expectedStatus: http.StatusBadRequest,
			expectedError:  gobperror.NewValidationError("name is a required field"),
		},
		"MinimumNameLength": {
			typ:            fdom.Chair,
			name:           "Ch",
			expectedStatus: http.StatusBadRequest,
			expectedError:  gobperror.NewValidationError("name must be at least 3 characters in length"),
		},
		"MaximumNameLength": {
			typ: fdom.Chair,
			name: "This is a gonna be a super lengthy name for a chair. " +
				"Trust me it really is gonna be so long that it won't fit into a single screen. " +
				"This is a gonna be a super lengthy name for a chair. Need more chars.",
			expectedStatus: http.StatusBadRequest,
			expectedError:  gobperror.NewValidationError("name must be a maximum of 200 characters in length"),
		},
	}

	for name, args := range subtests {
		t.Run(name, func(tt *testing.T) {
			body := fdom.AddFurnitureRequest{
				Type: args.typ,
				Name: args.name,
			}

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(body)
			if err != nil {
				tt.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/v1/furniture", &buf)
			if err != nil {
				tt.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			server, _ := testrest.SetupRESTServer(tt)
			handler := http.HandlerFunc(server.AddFurniture)

			handler.ServeHTTP(rr, req)
			assert.Equal(tt, args.expectedStatus, rr.Code)

			var response gobperror.GobpError
			err = json.Unmarshal(rr.Body.Bytes(), &response)

			assert.NoError(tt, err)
			assert.Equal(tt, args.expectedError.Code, response.Code)
			assert.Equal(tt, args.expectedError.Err, response.Err)
		})
	}
}

func TestAddFurnitureSuccess(t *testing.T) {
	body := fdom.AddFurnitureRequest{
		Type: fdom.Chair,
		Name: "New Fancy Chair",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/v1/furniture", &buf)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	server, useCases := testrest.SetupRESTServer(t)
	furnitureUseCase := useCases.Furniture.(*ucmock.FurnitureUseCaseInterface)
	furnitureUseCase.
		On("AddFurniture", mock.Anything, mock.Anything).
		Return(&fdom.FurnitureID{ID: 2}, nil).
		Once()
	handler := http.HandlerFunc(server.AddFurniture)

	// make http call
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}
