package handlers

import (
	"net/http"
	"service-media/helpers"
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/services"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandler interface {
	Create(ctx *gin.Context)
	MySocialMedia(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
		return
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success Social Media is Found", getSocmed)
}

func (h *SocialMediaHandlerImpl) Update(ctx *gin.Context) {
	var socmedInput web.SocialMediaRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&socmedInput)
	} else {
		ctx.ShouldBind(&socmedInput)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMediaId := ctx.Param("socialMediaId")
	socialMediaParse, _ := strconv.ParseUint(socialMediaId, 10, 32)

	updateUser, err := h.SocmedServices.Update(socmedInput, userId, uint(socialMediaParse))

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "Failed Update Social Media", err.Error())
		return
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success Updated Social Media", updateUser)

}

func (h *SocialMediaHandlerImpl) Delete(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	socialMediaParse, _ := strconv.ParseUint(socialMediaId, 10, 32)

	err := h.SocmedServices.Delete(uint(socialMediaParse))

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Social Media Not Found", err.Error())
	}

	convertBodyStatusResponse(ctx, http.StatusAccepted, "Deleted Success Social Media", "Deleted Success Social Media")
}
