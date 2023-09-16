package model

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID        int            `gorm:"type:bigint;primaryKey;autoIncrement"`
	Quantity  int            `gorm:"type:int"`
	Total     int            `gorm:"type:int"`
	ProductID int            `gorm:"type:bigint;not null"`
	UserID    int            `gorm:"type:bigint;not null"`
	Product   Products       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index"`
}
