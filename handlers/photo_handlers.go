package handlers

import (
	"fmt"
	"net/http"
	"service-media/helpers"
	"service-media/models/web"
	"service-media/services"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	Create(ctx *gin.Context)
	GetAllPhotos(ctx *gin.Context)
	GetPhotosByUser(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type PhotoHandlerImpl struct {
	PhotoService services.PhotoService
}

func NewPhotoHandler(photoService services.PhotoService) PhotoHandler {
	return &PhotoHandlerImpl{PhotoService: photoService}
}

func convertBodyStatusResponse(ctx *gin.Context, code int, message string, data interface{}) {
	switch code != 0 {
	case code == 201:
		ctx.JSON(code, gin.H{
			"message": message,
			"data":    data,
		})
	case code == 200:
		ctx.JSON(code, gin.H{
			"message": message,
			"data":    data,
		})
	default:
		ctx.JSON(code, gin.H{
			"message": message,
		})
	}
}

func (h *PhotoHandlerImpl) Create(ctx *gin.Context) {
	var PhotoInput web.PhotoRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&PhotoInput)
	} else {
		ctx.ShouldBind(&PhotoInput)
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	newPhoto, err := h.PhotoService.Create(PhotoInput, userID)

	fmt.Println(newPhoto.User)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "Failed created photo", err.Error())
		return
	}

	convertBodyStatusResponse(ctx, http.StatusCreated, "Success Created Photo", newPhoto)
}

func (h *PhotoHandlerImpl) GetAllPhotos(ctx *gin.Context) {
	photos, err := h.PhotoService.GetAllPhotos()

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}
	convertBodyStatusResponse(ctx, http.StatusOK, "All Photos", photos)
}

func (h *PhotoHandlerImpl) GetPhotosByUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	photos, err := h.PhotoService.GetPhotosByUser(userId)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Not Found Photos User", err.Error())
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success found photos user", photos)
}

func (h *PhotoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	var photoInput web.PhotoRequest

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&photoInput)
	} else {
		ctx.ShouldBind(&photoInput)
	}
	photoId := ctx.Param("photoId")
	parsePhotoID, _ := strconv.ParseUint(photoId, 10, 32)
	photo, err := h.PhotoService.UpdatePhoto(photoInput, uint(parsePhotoID))

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "Failed Update Photo", err.Error())
		return
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success Updated Photo", photo)
}

func (h *PhotoHandlerImpl) DeletePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	parsePhotoID, _ := strconv.ParseUint(photoId, 10, 32)
	err := h.PhotoService.DeletePhoto(uint(parsePhotoID))

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Failed Delete Photo", err.Error())
	}

	convertBodyStatusResponse(ctx, http.StatusAccepted, "Success Deleted Photo", "Success Deleted Photo")
}
