package cdom

import "gorm.io/gorm"

// It is not recommended to use basic types as key in context.WithValue().
// So we use ContextType for all context keys.
type ContextType string

const (
	Source ContextType = "Source"
)

type TxnManagerInterface interface {
	// ideally, core should should not be aware of gorm, but here we are
	GetTx() *gorm.DB
	Begin() (TxnManagerInterface, func())
	Commit() error
	Rollback() error
}
