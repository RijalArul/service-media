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

type CommentHandler interface {
	Create(ctx *gin.Context)
	GetComments(ctx *gin.Context)
}

type CommentHandlerImpl struct {
	CommentService services.CommentService
}

func NewCommentHandler(commentService services.CommentService) CommentHandler {
	return &CommentHandlerImpl{CommentService: commentService}
}

func (h *CommentHandlerImpl) Create(ctx *gin.Context) {
	var commentInput web.CommentRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&commentInput)
	} else {
		ctx.ShouldBind(&commentInput)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photoId := ctx.Param("photoId")
	parsePhotoID, _ := strconv.ParseUint(photoId, 10, 32)

	newComment, err := h.CommentService.Create(commentInput, userID, uint(parsePhotoID))

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadGateway, "Failed Comment", err.Error())
	}

	convertBodyStatusResponse(ctx, http.StatusCreated, "Success Created Comment", newComment)

}

func (h *CommentHandlerImpl) GetComments(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	getComments, err := h.CommentService.GetComments(userID)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Failed Not Found Comments", err.Error())
	}

	convertBodyStatusResponse(ctx, http.StatusOK, "Success Get Comments", getComments)
}
