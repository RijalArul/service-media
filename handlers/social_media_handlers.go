package handlers

import (
	"net/http"
	"service-media/helpers"
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandler interface {
	Create(ctx *gin.Context)
	MySocialMedia(ctx *gin.Context)
}

type SocialMediaHandlerImpl struct {
	SocmedServices services.SocialMediaService
}

func NewSocialMediaHandler(socmedService services.SocialMediaService) SocialMediaHandler {
	return &SocialMediaHandlerImpl{SocmedServices: socmedService}
}

func (h *SocialMediaHandlerImpl) Create(ctx *gin.Context) {
	var inputSocmed web.SocialMediaRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputSocmed)
	} else {
		ctx.ShouldBind(&inputSocmed)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	newSocmed, err := h.SocmedServices.Create(inputSocmed, userId)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "Failed Create Social Media", err.Error())
		return
	}

	convertBodyStatusResponse(ctx, http.StatusCreated, "Success Created Social Media", newSocmed)
}

func (h *SocialMediaHandlerImpl) MySocialMedia(ctx *gin.Context) {
	var socialmedias models.SocialMedia
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	getSocmed, err := h.SocmedServices.MySocialMedia(socialmedias, userId)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Failed Social Media Not Found", err.Error())
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success Social Media is Found", getSocmed)
}
