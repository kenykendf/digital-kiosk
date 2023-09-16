package repository

import (
	"kenykendf/digital-kiosk/internal/app/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WishlistRepo struct {
	DB *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) *WishlistRepo {
	return &WishlistRepo{DB: db}
}

func (wr *WishlistRepo) GetWishlistLists() ([]model.Wishlist, error) {
	var (
		wishlists []model.Wishlist
	)

	err := wr.DB.Debug().
		Preload("User").
		Preload("Products").
		Find(&wishlists).Error
	if err != nil {
		return wishlists, err
	}

	return wishlists, nil
}

func (wr *WishlistRepo) CreateWishlist(params model.Wishlist) error {
	wishlist := model.Wishlist{
		UserID:    params.UserID,
		ProductID: params.ProductID,
	}

	err := wr.DB.Create(&wishlist).Error
	if err != nil {
		logrus.Error("unable to create wishlist")
		return err
	}

	return nil
}

func (wr *WishlistRepo) DeleteWishlist(ID int) error {
	wishlist := &model.Wishlist{
		ID: uint(ID),
	}

	err := wr.DB.Delete(&wishlist).Error
	if err != nil {
		logrus.Error("unable to soft delete wishlist")
		return err
	}

	return nil
}
