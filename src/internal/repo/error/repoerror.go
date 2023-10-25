package repoerror

import "errors"

var (
	ErrNoDbConn     = errors.New("no database connection")
	ErrUnexpectedDB = errors.New("unexpected database error")
)
