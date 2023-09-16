package service

import (
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"
)

type ProductRepo interface {
	GetProductsLists(search schema.QueryParams) ([]model.Products, error)
	GetProductByID(ID string) (model.Products, error)
	CreateProduct(params model.Products) error
	UpdateProduct(ID string, params model.Products) error
	UpdateProductViews(ID string, params model.Products) error
	UpdateProductSold(ID string, params model.Products) error
	DeleteProduct(ID string) error
}

type ProductService struct {
	productRepo         ProductRepo
	productCategoryRepo ProductCategoryRepo
}

func NewProductService(productRepo ProductRepo, productCategoryRepo ProductCategoryRepo) *ProductService {
	return &ProductService{
		productRepo:         productRepo,
		productCategoryRepo: productCategoryRepo,
	}
}

func (ps *ProductService) GetProductsLists(search schema.QueryParams) ([]schema.GetProductsLists, error) {
	var response []schema.GetProductsLists

	data, err := ps.productRepo.GetProductsLists(search)
	if err != nil {
		return response, errors.New("unable to get categories lists")
	}

	for _, v := range data {
		var products schema.GetProductsLists
		products.ID = v.ID
		products.Name = v.Name
		products.Description = v.Description
		products.ProductCategoryID = v.ProductCategoryID
		products.ProductCategoryName = v.ProductCategory.Name
		products.Currency = v.Currency
		products.Price = v.Price
		products.Views = v.Views
		products.Sold = v.Sold
		products.Quantity = v.Quantity

		response = append(response, products)
	}

	return response, nil
}

func (ps *ProductService) GetProductByID(ID string) (schema.GetProductsLists, error) {
	var response schema.GetProductsLists

	data, err := ps.productRepo.GetProductByID(ID)
	if err != nil {
		return response, errors.New("unable to get product detail")
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description
	response.ProductCategoryID = data.ProductCategoryID
	response.ProductCategoryName = data.ProductCategory.Name
	response.Currency = data.Currency
	response.Price = data.Price
	response.Views = data.Views
	response.Sold = data.Sold
	response.Quantity = data.Quantity

	update := model.Products{}
	update.Views = data.Views + 1

	err = ps.productRepo.UpdateProduct(ID, update)
	if err != nil {
		return response, errors.New("increment failed")
	}

	return response, nil
}

func (ps *ProductService) CreateProduct(UserID int, params *schema.CreateProduct) error {
	id := strconv.FormatUint(uint64(params.ProductCategoryID), 10)

	_, err := ps.productCategoryRepo.GetProductCategoryByID(id)
	if err != nil {
		return errors.New("unable to get product detail")
	}

	product := model.Products{
		Name:              params.Name,
		Description:       params.Description,
		ProductCategoryID: params.ProductCategoryID,
		Currency:          params.Currency,
		Price:             params.Price,
		Quantity:          params.Quantity,
		UserID:            UserID,
	}

	err = ps.productRepo.CreateProduct(product)
	if err != nil {
		logrus.Error(reason.CreateProductErr)
		return err
	}

	return nil
}

func (ps *ProductService) UpdateProduct(ID string, params *schema.UpdateProduct) error {
	product := model.Products{}

	data, err := ps.GetProductByID(ID)
	if err != nil {
		return errors.New("unable to get product detail")
	}

	product.Name = params.Name
	if params.Name == "" {
		product.Name = data.Name
	}
	product.Description = params.Description
	if params.Description == "" {
		product.Description = data.Description
	}
	product.ProductCategoryID = params.ProductCategoryID
	if params.ProductCategoryID == 0 {
		product.ProductCategoryID = data.ProductCategoryID
	}
	product.Price = params.Price
	if params.Price == 0 {
		product.Price = data.Price
	}
	product.Currency = params.Currency
	if params.Currency == "" {
		product.Currency = data.Currency
	}
	product.Quantity = params.Quantity
	if params.Quantity == 0 {
		product.Quantity = data.Quantity
	}

	err = ps.productRepo.UpdateProduct(ID, product)
	if err != nil {
		logrus.Error("update product by ID : %w", err)
		return errors.New(reason.UpdateProductErr)
	}

	return nil
}

func (ps *ProductService) UpdateProductSell(ID string, params *schema.UpdateProductSell) error {
	product := model.Products{}

	data, err := ps.productRepo.GetProductByID(ID)
	if err != nil {
		return errors.New("unable to get product detail")
	}

	if data.Quantity < 1 {
		return errors.New("out of stock")
	}

	product.Sold = data.Sold + params.Quantity
	product.Quantity = data.Quantity - params.Quantity

	err = ps.productRepo.UpdateProduct(ID, product)
	if err != nil {
		logrus.Error("update product by ID : %w", err)
		return errors.New(reason.UpdateProductErr)
	}

	return nil
}

func (ps *ProductService) DeleteProduct(ID string) error {
	_, err := ps.GetProductByID(ID)
	if err != nil {
		return errors.New("unable to get product detail")
	}

	err = ps.productRepo.DeleteProduct(ID)
	if err != nil {
		logrus.Error("delete product by ID : %w", err)
		return errors.New(reason.DeleteProdcutErr)
	}

	return nil

}
