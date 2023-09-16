package controller

import (
	"net/http"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductCategoryService interface {
	GetProductCategoriesLists() ([]schema.GetProductCategoriesLists, error)
	GetProductCategoryByID(ID string) (schema.GetProductCategoriesLists, error)
	CreateProductCategory(params *schema.CreateProductCategory) error
	UpdateProductCategory(ID string, params schema.UpdateProductCategory) error
	DeleteProductCategory(ID string) error
}

type ProductCategoryController struct {
	service ProductCategoryService
}

func NewProductCategoryController(service ProductCategoryService) *ProductCategoryController {
	return &ProductCategoryController{service: service}
}

func (pc *ProductCategoryController) GetProductCategoriesLists(ctx *gin.Context) {
	data, err := pc.service.GetProductCategoriesLists()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.GetListsProductCatErr)
		return
	}
	if data == nil {
		data = make([]schema.GetProductCategoriesLists, 0)
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (pc *ProductCategoryController) GetProductCategoryByID(ctx *gin.Context) {
	productCategoryID := ctx.Param("id")

	data, err := pc.service.GetProductCategoryByID(productCategoryID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.GetDetailProductErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}
func (pc *ProductCategoryController) CreateProductCategory(ctx *gin.Context) {
	reqBody := schema.CreateProductCategory{}
	if handler.BindAndCheck(ctx, &reqBody) {
		return
	}

	err := pc.service.CreateProductCategory(&reqBody)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CreateProductCatErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (pc *ProductCategoryController) UpdateProductCategory(ctx *gin.Context) {
	productCategoryID := ctx.Param("id")

	reqBody := schema.UpdateProductCategory{}
	if handler.BindAndCheck(ctx, &reqBody) {
		logrus.Error(reason.UpdateProductCatErr)
		return
	}

	err := pc.service.UpdateProductCategory(productCategoryID, reqBody)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.UpdateProductCatErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)

}

func (pc *ProductCategoryController) DeleteProductCategory(ctx *gin.Context) {
	productCategoryID := ctx.Param("id")

	err := pc.service.DeleteProductCategory(productCategoryID)
	if err != nil {
		logrus.Error(reason.DeleteProdcutCatErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
