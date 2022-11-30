package handlers

import (
	"fmt"
	"net/http"
	"service-media/helpers"
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/services"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

func convertUserBodyResponse(user models.User) web.CreateUserResponse {
	return web.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func convertReturnUserResponse(ctx *gin.Context, code int, message string, data interface{}) {
	fmt.Println(code)
	switch code != 0 {
	case code == 201:
		ctx.JSON(code, gin.H{
			"message": message,
			"user":    data,
		})
	case code == 200:
		ctx.JSON(code, gin.H{
			"message": message,
			"token":   data,
		})
	default:
		ctx.JSON(code, gin.H{
			"message": message,
		})
	}
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
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "Bad Request", err.Error())

		return
	}

	userResp := convertUserBodyResponse(user)

	convertReturnUserResponse(ctx, http.StatusCreated, "Created User", userResp)

}

func (h *UserHandlerImpl) Login(ctx *gin.Context) {
	var LoginInput web.LoginUserRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&LoginInput)
	} else {
		ctx.ShouldBind(&LoginInput)
	}

	user, err := h.UserService.Login(LoginInput)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "User Not Found", err.Error())
		return
	}

	validPass := helpers.ComparePass([]byte(user.Password), []byte(LoginInput.Password))
	if !validPass {
		helpers.ConvertErrResponse(ctx, http.StatusUnauthorized, "Password Failed", "Password Failed")
		return
	}

	genToken := helpers.GenerateToken(user.ID, user.Email)

	convertReturnUserResponse(ctx, http.StatusOK, "Login Success", genToken)
}
