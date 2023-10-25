package resth

import (
	"encoding/json"
	"io"
	"net/http"

	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	"github.com/mrpandey/gobp/src/util"
)

func SendErrorResponse(logger *util.StandardLogger, w http.ResponseWriter, err error) {
	switch errPayload := err.(type) { //nolint:errorlint
	case gobperror.GobpError:
		sendJSONResponse(logger, w, errPayload.HttpStatus, errPayload)
	default:
		errUnexpected := gobperror.ErrUnexpected
		sendJSONResponse(logger, w, errUnexpected.HttpStatus, errUnexpected)
	}
}

func sendJSONResponse(logger *util.StandardLogger, w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// send empty body instead of 'null' if payload is nil
	if payload == nil {
		return
	}

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {

		logger.Errorf("Error sending JSON response %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func decodeJSON(logger *util.StandardLogger, body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&v)
	if err != nil {
		switch err := err.(type) { //nolint:errorlint
		case *json.UnmarshalTypeError:
			return gobperror.NewInvalidFieldTypeError(err.Field)
		case *json.SyntaxError:
			return gobperror.NewJsonSyntaxError(err.Error())
		default:

			logger.Errorf("Unexpected error while decoding JSON. %v", err)
			return gobperror.ErrJsonDecode
		}
	}
	return nil
}
