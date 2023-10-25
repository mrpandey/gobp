package gobperror

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// GobpError implements error interface.
type GobpError struct {
	Code       string     `json:"code"`
	Err        string     `json:"error"`
	HttpStatus int        `json:"-"`
	GrpcCode   codes.Code `json:"-"`
}

func (e GobpError) Error() string {
	return e.Code + ": " + e.Err
}

var ErrUnexpected = GobpError{
	Code:       "E000",
	Err:        "Unexpected error.",
	HttpStatus: http.StatusInternalServerError,
	GrpcCode:   codes.Internal,
}

func NewJsonSyntaxError(msg string) GobpError {
	return GobpError{
		Code:       "E001",
		Err:        "Invalid JSON syntax: " + msg,
		HttpStatus: http.StatusBadRequest,
		GrpcCode:   codes.InvalidArgument,
	}
}

var ErrJsonDecode = GobpError{
	Code:       "E002",
	Err:        "Could not decode JSON in request body. Check your payload.",
	HttpStatus: http.StatusBadRequest,
	GrpcCode:   codes.InvalidArgument,
}

func NewValidationError(err string) GobpError {
	return GobpError{
		Code:       "E003",
		Err:        err,
		HttpStatus: http.StatusBadRequest,
		GrpcCode:   codes.InvalidArgument,
	}
}

func NewDuplicateKeyError(entityName string, keyName string) GobpError {
	return GobpError{
		Code:       "E004",
		Err:        entityName + " with provided " + keyName + " already exists.",
		HttpStatus: http.StatusBadRequest,
		GrpcCode:   codes.AlreadyExists,
	}
}

var ErrRecordNotFound = GobpError{
	Code:       "E005",
	Err:        "Record not found.",
	HttpStatus: http.StatusNotFound,
	GrpcCode:   codes.NotFound,
}

var ErrBlocked = GobpError{
	Code:       "E006",
	Err:        "The client is blocked.",
	HttpStatus: http.StatusNotFound,
	GrpcCode:   codes.NotFound,
}

var ErrAuthHeaderMissing = GobpError{
	Code:       "E007",
	Err:        "Authorization missing or malformed.",
	HttpStatus: http.StatusUnauthorized,
	GrpcCode:   codes.Unauthenticated,
}

var ErrWrongCreds = GobpError{
	Code:       "E008",
	Err:        "Provided credentials are wrong.",
	HttpStatus: http.StatusUnauthorized,
	GrpcCode:   codes.Unauthenticated,
}

var ErrInvalidToken = GobpError{
	Code:       "E009",
	Err:        "Provided token is invalid or expired.",
	HttpStatus: http.StatusUnauthorized,
	GrpcCode:   codes.Unauthenticated,
}

func NewInvalidFieldTypeError(fieldName string) GobpError {
	errMsg := fmt.Sprintf("Wrong type provided for field '%v'.", fieldName)
	return NewValidationError(errMsg)
}

func NewInvalidUrlParamTypeError(paramName string) GobpError {
	errMsg := fmt.Sprintf("Wrong type provided for URL parameter '%v'.", paramName)
	return NewValidationError(errMsg)
}
