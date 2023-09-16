package service

import (
	"errors"
	"fmt"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"github.com/sirupsen/logrus"
)

type ProductCategoryRepo interface {
	GetProductCategoriesLists() ([]model.ProductCategories, error)
	GetProductCategoryByID(ID string) (model.ProductCategories, error)
	CreateProductCategory(params *model.ProductCategories) error
	UpdateProductCategory(ID string, params model.ProductCategories) error
	DeleteProductCategory(ID string) error
}

type ProductCategoryService struct {
	repo ProductCategoryRepo
}

func NewProductCategoryService(repo ProductCategoryRepo) *ProductCategoryService {
	return &ProductCategoryService{repo: repo}
}

func (pcs *ProductCategoryService) GetProductCategoriesLists() ([]schema.GetProductCategoriesLists, error) {
	var response []schema.GetProductCategoriesLists

	data, err := pcs.repo.GetProductCategoriesLists()
	if err != nil {
		return response, errors.New(reason.GetListsProductCatErr)
	}

	for _, v := range data {
		var products schema.GetProductCategoriesLists
		products.ID = v.ID
		products.Name = v.Name
		products.Description = v.Description

		response = append(response, products)
	}

	return response, nil
}

func (pcs *ProductCategoryService) GetProductCategoryByID(ID string) (schema.GetProductCategoriesLists, error) {
	var response schema.GetProductCategoriesLists

	data, err := pcs.repo.GetProductCategoryByID(ID)
	if err != nil {
		return response, errors.New(reason.GetDetailProductCatErr)
	}
	if data.ID < 1 {
		logrus.Error(fmt.Errorf("category ID not found"))
		return schema.GetProductCategoriesLists{}, nil
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description

	return response, nil
}

func (pcs *ProductCategoryService) CreateProductCategory(params *schema.CreateProductCategory) error {
	productCategory := model.ProductCategories{
		Name:        params.Name,
		Description: params.Description,
	}

	err := pcs.repo.CreateProductCategory(&productCategory)
	if err != nil {
		logrus.Error(reason.CreateProductCatErr)
		return err
	}

	return nil
}

func (pcs *ProductCategoryService) UpdateProductCategory(ID string, params schema.UpdateProductCategory) error {
	var update model.ProductCategories

	productCategory, err := pcs.repo.GetProductCategoryByID(ID)
	if err != nil {
		logrus.Error(reason.GetDetailProductCatErr)
		return err
	}

	update.Name = params.Name
	if params.Name == "" {
		update.Name = productCategory.Name
	}

	update.Description = params.Description
	if params.Description == "" {
		update.Description = productCategory.Description
	}

	err = pcs.repo.UpdateProductCategory(ID, update)
	if err != nil {
		logrus.Error(reason.UpdateProductCatErr)
		return err
	}

	return nil
}

func (pcs *ProductCategoryService) DeleteProductCategory(ID string) error {
	_, err := pcs.repo.GetProductCategoryByID(ID)
	if err != nil {
		logrus.Error(reason.GetDetailProductCatErr)
		return err
	}

	err = pcs.repo.DeleteProductCategory(ID)
	if err != nil {
		logrus.Error(reason.DeleteProdcutCatErr)
		return err
	}

	return nil
}
