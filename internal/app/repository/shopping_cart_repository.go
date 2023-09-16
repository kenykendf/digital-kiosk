package repository

import (
	"fmt"
	"time"

	"kenykendf/digital-kiosk/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	DB *gorm.DB
}

func NewShoppingCartRepository(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{DB: db}
}

func (scr *ShoppingCartRepository) Create(ShoppingCart model.ShoppingCart) error {

	if err := scr.DB.Omit("updated_at", "deleted_at").Create(&ShoppingCart).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingCartRepository - Create : %w", err))
		return err
	}
	return nil
}

func (scr *ShoppingCartRepository) CheckProduct(UserID int, ProductID int) (bool, error) {
	var count int64

	if err := scr.DB.Model(&model.ShoppingCart{}).Where("user_id = ? AND product_id = ?", UserID, ProductID).Count(&count).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingCartRepository - CheckProduct : %w", err))
		return false, err
	}

	return count > 0, nil
}

func (scr *ShoppingCartRepository) GetByID(id string) (model.ShoppingCart, error) {

	var ShoppingCart model.ShoppingCart

	result := scr.DB.Find(&ShoppingCart, id)
	if result.Error != nil {
		log.Error(fmt.Errorf("error ShoppingCartRepository - GetByID : %w", result.Error))
		return ShoppingCart, result.Error
	}

	return ShoppingCart, nil
}

func (scr *ShoppingCartRepository) Browse(UserID int) ([]model.ShoppingCart, error) {
	var ShoppingCarts []model.ShoppingCart

	result := scr.DB.Preload("Product").Where("user_id", UserID).Find(&ShoppingCarts)
	if result.Error != nil {
		log.Error(fmt.Errorf("error ShoppingCartRepository - Browse : %w", result.Error))
		return ShoppingCarts, result.Error

	}
	return ShoppingCarts, nil
}

func (scr *ShoppingCartRepository) UpdateByID(id string, shoppingCart model.ShoppingCart) error {

	data := &model.ShoppingCart{
		Quantity:  shoppingCart.Quantity,
		Total:     shoppingCart.Total,
		UpdatedAt: time.Now(),
	}

	if err := scr.DB.Where("id = ?", id).Updates(&data).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingCartRepository - UpdateByID : %w", err))
		return err
	}

	return nil
}

func (scr *ShoppingCartRepository) DeleteByID(id string) error {
	var ShoppingCart model.ShoppingCart

	if err := scr.DB.Where("id = ?", id).Delete(&ShoppingCart).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingCartRepository - DeleteByID : %w", err))
		return err
	}

	return nil
}
