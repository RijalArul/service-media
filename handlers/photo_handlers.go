package handlers

import (
	"net/http"
	"service-media/helpers"
	"service-media/models/web"
	"service-media/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	Create(ctx *gin.Context)
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
	default:
		ctx.JSON(code, message)
	}
}

func (h *PhotoHandlerImpl) Create(ctx *gin.Context) {
	var PhotoInput web.CreaatePhotoRequest
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
	}

	convertBodyStatusResponse(ctx, http.StatusCreated, "Success Created Photo", newPhoto)
}
