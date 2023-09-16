package controller

import (
	"net/http"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"

	"github.com/gin-gonic/gin"
)

type ShoppingCartService interface {
	Create(req *schema.CreateShoppingCartReq) error
	BrowseAll(req *schema.GetShoppingCartReq) ([]schema.GetShoppingCartResp, error)
	Update(id string, req *schema.UpdateShoppingCartReq) error
	Delete(id string) error
}

type ShoppingCartController struct {
	cartService    ShoppingCartService
	productService ProductService
}

func NewShoppingCartController(cartService ShoppingCartService, productService ProductService) *ShoppingCartController {
	return &ShoppingCartController{cartService: cartService, productService: productService}
}

func (scc *ShoppingCartController) CreateShoppingCart(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.CreateShoppingCartReq{}
	req.UserID = id

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := scc.cartService.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create Shopping Cart", nil)
}

func (scc *ShoppingCartController) BrowseShoppingCart(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.GetShoppingCartReq{}

	req.UserID = userID

	resp, err := scc.cartService.BrowseAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if resp == nil {
		resp = make([]schema.GetShoppingCartResp, 0)
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get list Shopping Cart", resp)
}

func (scc *ShoppingCartController) UpdateShoppingCart(ctx *gin.Context) {

	req := &schema.UpdateShoppingCartReq{}
	id, _ := ctx.Params.Get("id")
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))

	req.UserID = userID

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := scc.cartService.Update(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update Shopping Cart", nil)
}

func (scc *ShoppingCartController) DeleteShoppingCart(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := scc.cartService.Delete(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete Shopping Cart", nil)
}
