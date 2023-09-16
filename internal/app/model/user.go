package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            int            `gorm:"type:bigint;primaryKey;autoIncrement"`
	Email         string         `gorm:"type:varchar(255);unique;not null"`
	Password      string         `gorm:"type:varchar(255);not null"`
	Fullname      string         `gorm:"type:varchar(255);not null"`
	CreatedAt     time.Time      `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt     time.Time      `gorm:"type:timestamp"`
	DeletedAt     gorm.DeletedAt `gorm:"type:timestamp;index"`
	Auths         Auth           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShoppingCarts ShoppingCart   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Products      Products       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
