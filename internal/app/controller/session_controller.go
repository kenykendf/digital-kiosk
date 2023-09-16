package controller

import (
	"net/http"
	"strconv"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"

	"github.com/gin-gonic/gin"
)

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
	RefreshToken(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
	Logout(req *schema.LogoutReq) error
}

type RefreshTokenVerifier interface {
	ValidateRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	service    SessionService
	tokenMaker RefreshTokenVerifier
}

func NewSessionController(service SessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{service: service, tokenMaker: tokenMaker}
}

func (sc *SessionController) Login(ctx *gin.Context) {
	req := &schema.LoginReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	token, err := sc.service.Login(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success login", token)
}

func (sc *SessionController) Refresh(ctx *gin.Context) {
	req := &schema.RefreshTokenReq{}
	refreshToken := ctx.GetHeader("refresh_token")

	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot refresh token")
		return
	}

	validate, err := sc.tokenMaker.ValidateRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	UserID, _ := strconv.Atoi(validate)

	req.UserID = UserID
	req.RefreshToken = refreshToken

	token, err := sc.service.RefreshToken(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh token", token)
}

func (sc *SessionController) Logout(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.LogoutReq{}

	req.UserID = id

	err := sc.service.Logout(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success logout", nil)
}
