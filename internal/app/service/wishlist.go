package service

import (
	"errors"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"github.com/sirupsen/logrus"
)

type WishlistRepo interface {
	GetWishlistLists() ([]model.Wishlist, error)
	CreateWishlist(params model.Wishlist) error
	DeleteWishlist(ID int) error
}

type WishlistService struct {
	wishlistRepo WishlistRepo
	productRepo  ProductRepo
}

func NewWishlistService(wishlistRepo WishlistRepo, productRepo ProductRepo) *WishlistService {
	return &WishlistService{
		wishlistRepo: wishlistRepo,
		productRepo:  productRepo,
	}
}

func (ws *WishlistService) GetWishlistLists() ([]schema.GetWishlists, error) {
	var response []schema.GetWishlists

	data, err := ws.wishlistRepo.GetWishlistLists()
	if err != nil {
		return nil, errors.New(reason.GetListsWishlistErr)
	}

	for _, v := range data {
		var wishlist schema.GetWishlists
		wishlist.ID = v.ID
		wishlist.Currency = v.Products.Currency
		wishlist.Description = v.Products.Description
		wishlist.Name = v.Products.Name
		wishlist.Price = v.Products.Price
		wishlist.ProductCategoryName = v.Products.ProductCategory.Name
		wishlist.ProductID = v.Products.ID
		wishlist.Quantity = v.Products.Quantity
		wishlist.Sold = v.Products.Sold
		wishlist.UserID = uint(v.UserID)

		response = append(response, wishlist)
	}

	return response, nil
}

func (ws *WishlistService) CreateWishlist(params *schema.CreateWishlist) error {
	productID := strconv.Itoa(params.ProductID)
	_, err := ws.productRepo.GetProductByID(productID)
	if err != nil {
		return errors.New("unable to get product detail")
	}

	wishlist := model.Wishlist{
		ProductID: params.ProductID,
		UserID:    params.UserID,
	}

	err = ws.wishlistRepo.CreateWishlist(wishlist)
	if err != nil {
		logrus.Error(reason.CreateWishlistErr)
		return err
	}

	return nil
}

func (ws *WishlistService) DeleteWishlist(ID int) error {
	err := ws.wishlistRepo.DeleteWishlist(ID)
	if err != nil {
		logrus.Error(reason.DeleteWishlistErr)
		return err
	}

	return nil
}
