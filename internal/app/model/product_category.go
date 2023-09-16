package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategories struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"column:name;unique"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
