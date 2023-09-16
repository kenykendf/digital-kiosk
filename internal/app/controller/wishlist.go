package controller

import (
	"net/http"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"github.com/gin-gonic/gin"
)

type WishlistService interface {
	GetWishlistLists() ([]schema.GetWishlists, error)
	CreateWishlist(params *schema.CreateWishlist) error
	DeleteWishlist(ID int) error
}

type WishlistController struct {
	wishlistService WishlistService
}

func NewWishlistController(wishlistService WishlistService) *WishlistController {
	return &WishlistController{wishlistService: wishlistService}
}

func (wc *WishlistController) GetWishlistLists(ctx *gin.Context) {
	data, err := wc.wishlistService.GetWishlistLists()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.GetListsWishlistErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (wc *WishlistController) CreateWishlist(ctx *gin.Context) {
	req := &schema.CreateWishlist{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := wc.wishlistService.CreateWishlist(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CreateWishlistErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "wishlist created", nil)
}

func (wc *WishlistController) DeleteWishlist(ctx *gin.Context) {
	wishlistID := ctx.Param("id")
	ID, _ := strconv.Atoi(wishlistID)

	err := wc.wishlistService.DeleteWishlist(ID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.DeleteWishlistErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "wishlist deleted", nil)
}
