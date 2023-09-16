package model

import (
	"gorm.io/gorm"
)

type Products struct {
	ID                uint              `gorm:"primaryKey"`
	Name              string            `gorm:"column:name;not null;size:256;index;"`
	Description       string            `gorm:"column:description;type:text;not null"`
	Views             uint64            `gorm:"column:views;comment:product being visited"`
	Sold              uint64            `gorm:"column:sold"`
	ProductCategoryID uint              `gorm:"column:product_category_id;comment:foreignKey of product categories"`
	ProductCategory   ProductCategories `gorm:"foreignKey:ProductCategoryID;references:ID"`
	Currency          string            `gorm:"column:currency;size:3"`
	Price             uint64            `gorm:"column:price"`
	Quantity          uint64            `gorm:"column:quantity"`
	UserID            int               `gorm:"column:user_id"`
	gorm.Model
}
