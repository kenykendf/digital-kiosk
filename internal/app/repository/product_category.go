package repository

import (
	"strconv"
	"time"

	"kenykendf/digital-kiosk/internal/app/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductCategoryRepo struct {
	DB *gorm.DB
}

func NewProductCategoryRepo(db *gorm.DB) *ProductCategoryRepo {
	return &ProductCategoryRepo{DB: db}
}

func (pcr *ProductCategoryRepo) GetProductCategoriesLists() ([]model.ProductCategories, error) {
	var (
		productCategories []model.ProductCategories
	)

	pcr.DB.Model(productCategories).Scan(&productCategories)

	return productCategories, nil
}

func (pcr *ProductCategoryRepo) GetProductCategoryByID(ID string) (model.ProductCategories, error) {
	var (
		productCategories model.ProductCategories
	)

	pcr.DB.Model(productCategories).Where("id = ?", ID).Scan(&productCategories)

	return productCategories, nil
}

func (pcr *ProductCategoryRepo) CreateProductCategory(params *model.ProductCategories) error {
	productCategory := model.ProductCategories{
		Name:        params.Name,
		Description: params.Description,
	}

	err := pcr.DB.Create(&productCategory).Error
	if err != nil {
		logrus.Error("unable to create product category")
		return err
	}

	return nil
}

func (pcr *ProductCategoryRepo) UpdateProductCategory(ID string, params model.ProductCategories) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logrus.Error("error parsing ID string : ", err)
		return err
	}

	productCategory := &model.ProductCategories{
		ID:          uint(id),
		Name:        params.Name,
		Description: params.Description,
		UpdatedAt:   time.Now(),
	}

	err = pcr.DB.Updates(&productCategory).Error
	if err != nil {
		logrus.Error("unable to update product category")
		return err
	}

	return nil
}

func (pcr *ProductCategoryRepo) DeleteProductCategory(ID string) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logrus.Error("error parsing ID string : ", err)
		return err
	}

	productCategory := &model.ProductCategories{
		ID: uint(id),
	}

	err = pcr.DB.Delete(&productCategory).Error
	if err != nil {
		logrus.Error("unable to soft delete product category")
		return err
	}

	return nil
}
