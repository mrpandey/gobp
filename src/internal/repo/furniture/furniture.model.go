package frepo

import (
	"time"

	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"

	"gorm.io/gorm"
)

type Furniture struct {
	ID        uint
	Type      fdom.FurnitureType `gorm:"not null"`
	Name      string             `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
