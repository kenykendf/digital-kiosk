package controller

import (
	"net/http"

	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/handler"

	"github.com/gin-gonic/gin"
)

type RegistrationService interface {
	Register(req *schema.RegisterReq) error
}

type RegistrationController struct {
	service RegistrationService
}

func NewRegistrationController(service RegistrationService) *RegistrationController {
	return &RegistrationController{service: service}
}

func (rc *RegistrationController) Register(ctx *gin.Context) {
	req := &schema.RegisterReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := rc.service.Register(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success register", nil)
}
