package controller

import (
	"net/http"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	GetProductsLists(search schema.QueryParams) ([]schema.GetProductsLists, error)
	GetProductByID(ID string) (schema.GetProductsLists, error)
	CreateProduct(UserID int, params *schema.CreateProduct) error
	UpdateProduct(ID string, params *schema.UpdateProduct) error
	UpdateProductSell(ID string, params *schema.UpdateProductSell) error
	DeleteProduct(ID string) error
}

type ProductController struct {
	service ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) GetProductsLists(ctx *gin.Context) {
	search := schema.QueryParams{}
	search.Limit = ctx.GetInt("limit")
	search.Offset = ctx.GetInt("offset")
	search.Name = ctx.Query("name")
	search.Category = ctx.Query("category")
	search.SortBy = ctx.Query("sort_by")
	search.AscDesc = ctx.GetInt("asc")

	data, err := pc.service.GetProductsLists(search)
	if err != nil {
		handler.ResponseSuccess(ctx, http.StatusUnprocessableEntity, reason.GetListsProductErr, nil)
		return
	}
	if data == nil {
		data = make([]schema.GetProductsLists, 0)
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (pc *ProductController) GetProductsListsByViews(ctx *gin.Context) {
	search := schema.QueryParams{}
	search.Limit = ctx.GetInt("limit")
	search.Offset = ctx.GetInt("offset")
	search.Name = ctx.Query("name")
	search.Category = ctx.Query("category")
	search.SortBy = "views"
	search.AscDesc = 0

	data, err := pc.service.GetProductsLists(search)
	if err != nil {
		handler.ResponseSuccess(ctx, http.StatusUnprocessableEntity, reason.GetListsProductErr, nil)
		return
	}
	if data == nil {
		data = make([]schema.GetProductsLists, 0)
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (pc *ProductController) GetProductsListsBySold(ctx *gin.Context) {
	search := schema.QueryParams{}
	search.Limit = ctx.GetInt("limit")
	search.Offset = ctx.GetInt("offset")
	search.Name = ctx.Query("name")
	search.Category = ctx.Query("category")
	search.SortBy = "sold"
	search.AscDesc = 0

	data, err := pc.service.GetProductsLists(search)
	if err != nil {
		handler.ResponseSuccess(ctx, http.StatusUnprocessableEntity, reason.GetListsProductErr, nil)
		return
	}
	if data == nil {
		data = make([]schema.GetProductsLists, 0)
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (pc *ProductController) GetProductByID(ctx *gin.Context) {
	productID := ctx.Param("id")

	data, err := pc.service.GetProductByID(productID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.GetDetailProductErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	id, _ := strconv.Atoi(userID)

	req := &schema.CreateProduct{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := pc.service.CreateProduct(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CreateProductErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &schema.UpdateProduct{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := pc.service.UpdateProduct(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.UpdateProductErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (pc *ProductController) UpdateProductSell(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &schema.UpdateProductSell{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := pc.service.UpdateProductSell(id, req)
	if err != nil {
		logrus.Println("Check err ", err)
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.UpdateProductErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	err := pc.service.DeleteProduct(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.DeleteProdcutErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
