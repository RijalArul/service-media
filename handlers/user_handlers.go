package handlers

import (
	"net/http"
	"service-media/helpers"
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/services"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

var (
	appJSON = "application/json"
)

func (h *UserHandlerImpl) Register(ctx *gin.Context) {
	var UserInput web.CreateUserRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&UserInput)
	} else {
		ctx.ShouldBind(&UserInput)
	}

	user, err := h.UserService.Create(UserInput)

	if err != nil {
		convertErrUserResponse(ctx, http.StatusBadRequest, "Bad Request", err.Error())

		return
	}

	userResp := convertUserBodyResponse(user)

	convertReturnUserResponse(ctx, http.StatusCreated, "Created User", userResp)

}

func convertUserBodyResponse(user models.User) web.CreateUserResponse {
	return web.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func convertReturnUserResponse(ctx *gin.Context, code int, message, data interface{}) {
	ctx.JSON(code, gin.H{
		"message": message,
		"user":    data,
	})
}

func convertErrUserResponse(ctx *gin.Context, code int, message, err interface{}) {
	ctx.JSON(code, gin.H{
		"message": "message",
		"err":     err,
	})
}
