package authrepo

import "time"

type ClientCred struct {
	ID           uint   `gorm:"primaryKey"`
	Slug         string `gorm:"not null;uniqueIndex"`
	HashedSecret string `gorm:"not null"`
	IsBlocked    bool   `gorm:"not null"`

	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
