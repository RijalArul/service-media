package handlers

import (
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
			"photo":   data,
		})
	case code == 200:
		ctx.JSON(code, gin.H{
			"message": message,
			"photos":  data,
		})
	default:
		ctx.JSON(code, message)
	}
}

func (h *PhotoHandlerImpl) Create(ctx *gin.Context) {
	var PhotoInput web.PhotoRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&PhotoInput)
	} else {
		ctx.ShouldBind(&PhotoInput)
	}

	newPhoto, err := h.PhotoService.Create(PhotoInput, userID)

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

	// if contentType == appJSON {
	// 	ctx.ShouldBindJSON(photoInput)
	// } else {
	// 	ctx.ShouldBind(photoInput)
	// }

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "Failed Update Photo", err.Error())
		return
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success Updated Photo", photo)
}
