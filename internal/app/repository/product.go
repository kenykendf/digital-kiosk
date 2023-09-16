package repository

import (
	"fmt"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (pr *ProductRepo) GetProductsLists(search schema.QueryParams) ([]model.Products, error) {
	var (
		products = []model.Products{}
		limit    = search.Limit
		offset   = search.Offset
		sort     = search.SortBy
		name     = search.Name
		category = search.Category
	)

	if search.AscDesc == 1 {
		sort = sort + " asc"
	} else {
		sort = sort + " desc"
	}

	if search.Name == "" || search.Category == "" {
		err := pr.DB.Debug().Preload("ProductCategory").
			Order(sort).
			Limit(limit).
			Offset(offset).
			Find(&products).Error
		if err != nil {
			return products, err
		}
	}

	if search.Name != "" || search.Category != "" {
		err := pr.DB.Debug().
			Preload("ProductCategory", "name ilike ?", "%"+category+"%").
			Order(sort).
			Limit(limit).
			Offset(offset).
			Where("name ilike ?", "%"+name+"%").
			Find(&products).Error
		if err != nil {
			return products, err
		}
	}

	return products, nil
}

func (pr *ProductRepo) GetProductByID(ID string) (model.Products, error) {
	var (
		product model.Products
	)

	err := pr.DB.Preload("ProductCategory").Model(&product).
		Where("id = ?", ID).Find(&product).Error

	if err != nil {
		logrus.Error("unable to retrieve product detail")
		return product, err
	}
	return product, nil
}

func (pr *ProductRepo) CreateProduct(params model.Products) error {
	product := model.Products{
		Name:              params.Name,
		Description:       params.Description,
		Views:             0,
		Sold:              0,
		ProductCategoryID: params.ProductCategoryID,
		Currency:          params.Currency,
		Price:             params.Price,
		UserID:            params.UserID,
		Quantity:          params.Quantity,
	}

	err := pr.DB.Create(&product).Error
	if err != nil {
		logrus.Error("unable to create product")
		return err
	}

	return nil
}

func (pr *ProductRepo) UpdateProduct(ID string, params model.Products) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logrus.Error("error parsing ID string : ", err)
		return err
	}

	product := &model.Products{
		ID:                uint(id),
		Name:              params.Name,
		Description:       params.Description,
		ProductCategoryID: params.ProductCategoryID,
		Price:             params.Price,
		Views:             params.Views,
		Sold:              params.Sold,
		Currency:          params.Currency,
		Quantity:          params.Quantity,
	}
	fmt.Println("VALUES : ", product.Quantity)

	err = pr.DB.Updates(&product).Error
	if err != nil {
		logrus.Error("unable to update product")
		return err
	}

	return nil
}

func (pr *ProductRepo) UpdateProductViews(ID string, params model.Products) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logrus.Error("error parsing ID string : ", err)
		return err
	}

	product := &model.Products{
		ID:    uint(id),
		Views: params.Views + 1,
	}

	err = pr.DB.Updates(&product).Error
	if err != nil {
		logrus.Error("unable to update product views")
		return err
	}

	return nil
}

func (pr *ProductRepo) UpdateProductSold(ID string, params model.Products) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logrus.Error("error parsing ID string : ", err)
		return err
	}

	product := &model.Products{
		ID:   uint(id),
		Sold: params.Sold + 1,
	}

	err = pr.DB.Updates(&product).Error
	if err != nil {
		logrus.Error("unable to update product sold")
		return err
	}

	return nil
}

func (pr *ProductRepo) DeleteProduct(ID string) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logrus.Error("error parsing ID string : ", err)
		return err
	}

	productCategory := &model.Products{
		ID: uint(id),
	}

	err = pr.DB.Delete(&productCategory).Error
	if err != nil {
		logrus.Error("unable to soft delete product")
		return err
	}

	return nil
}
