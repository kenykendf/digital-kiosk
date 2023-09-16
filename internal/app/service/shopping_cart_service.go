package service

import (
	"errors"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"
)

type ShoppingCartService struct {
	cartRepo    ShoppingCartRepository
	productRepo ProductRepo
}

type ShoppingCartRepository interface {
	Create(ShoppingCart model.ShoppingCart) error
	CheckProduct(UserID int, ProductID int) (bool, error)
	GetByID(id string) (model.ShoppingCart, error)
	Browse(UserID int) ([]model.ShoppingCart, error)
	UpdateByID(id string, shoppingCart model.ShoppingCart) error
	DeleteByID(id string) error
}

func NewShoppingCartService(cartRepo ShoppingCartRepository, productRepo ProductRepo) *ShoppingCartService {
	return &ShoppingCartService{cartRepo: cartRepo, productRepo: productRepo}
}

func (scs *ShoppingCartService) Create(req *schema.CreateShoppingCartReq) error {
	var insertData model.ShoppingCart
	id := strconv.Itoa(req.ProductID)

	exists, err := scs.cartRepo.CheckProduct(req.UserID, req.ProductID)
	if err != nil {
		return errors.New("failed to check product existence")
	}
	if exists {
		return errors.New("product already exist in cart")
	}

	product, _ := scs.productRepo.GetProductByID(id)
	if product.ID == 0 {
		return errors.New("unable to get product detail")
	}

	insertData.Quantity = req.Quantity
	insertData.Total = int(product.Price) * req.Quantity
	insertData.UserID = req.UserID
	insertData.ProductID = req.ProductID

	err = scs.cartRepo.Create(insertData)
	if err != nil {
		return errors.New(reason.ShoppingCartCannotCreate)
	}

	return nil

}

func (scs *ShoppingCartService) BrowseAll(req *schema.GetShoppingCartReq) ([]schema.GetShoppingCartResp, error) {
	var resp []schema.GetShoppingCartResp

	Carts, err := scs.cartRepo.Browse(req.UserID)
	if err != nil {
		return nil, errors.New(reason.ShoppingCartCannotBrowse)
	}

	for _, value := range Carts {
		var respData schema.GetShoppingCartResp
		respData.ID = value.ID
		respData.Quantity = value.Quantity
		respData.Total = value.Total
		respData.Product.ProductID = int(value.Product.ID)
		respData.Product.ProductName = value.Product.Name
		respData.Product.ProductPrice = value.Product.Price

		resp = append(resp, respData)
	}

	return resp, nil
}

func (scs *ShoppingCartService) Update(id string, req *schema.UpdateShoppingCartReq) error {
	var updateData model.ShoppingCart

	cart, err := scs.cartRepo.GetByID(id)
	if err != nil {
		return errors.New(reason.ShoppingCartCannotGetDetail)
	}
	if cart.ID == 0 {
		return errors.New(reason.ShoppingCartNotFound)
	}
	productID := strconv.Itoa(cart.ProductID)

	product, _ := scs.productRepo.GetProductByID(productID)
	if product.ID == 0 {
		return errors.New("unable to get product detail")
	}

	updateData.Quantity = req.Quantity
	updateData.Total = int(product.Price) * req.Quantity

	err = scs.cartRepo.UpdateByID(id, updateData)
	if err != nil {
		return errors.New(reason.ShoppingCartCannotUpdate)
	}

	return nil

}

func (scs *ShoppingCartService) Delete(id string) error {

	check, err := scs.cartRepo.GetByID(id)
	if check.ID == 0 || err != nil {
		return errors.New(reason.ShoppingCartNotFound)
	}

	err = scs.cartRepo.DeleteByID(id)
	if err != nil {
		return errors.New(reason.ShoppingCartCannotDelete)
	}

	return nil
}
