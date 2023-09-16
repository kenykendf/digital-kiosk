package controller

import (
	"net/http"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	BrowseAll() ([]schema.GetUsersResp, error)
	GetByID(id int) (schema.GetUsersResp, error)
	DeleteByID(id int) error
}

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) BrowseUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.GetUserReq{}

	req.UserID = id

	users, err := uc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get list user ", users)
}

func (uc *UserController) DetailUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.GetUserReq{}

	req.UserID = userID

	resp, err := uc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get detail user", resp)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.GetUserReq{}

	req.UserID = userID

	err := uc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete user", nil)
}
