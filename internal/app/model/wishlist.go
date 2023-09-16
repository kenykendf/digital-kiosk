package model

import "gorm.io/gorm"

type Wishlist struct {
	ID        uint     `gorm:"primaryKey"`
	UserID    int      `gorm:"column:user_id;uniqueIndex:idx_wishlist"`
	ProductID int      `gorm:"column:product_id;uniqueIndex:idx_wishlist"`
	Products  Products `gorm:"foreignKey:ProductID;References:ID;"`
	User      User     `gorm:"foreignKey:UserID;References:ID;"`
	gorm.Model
}
