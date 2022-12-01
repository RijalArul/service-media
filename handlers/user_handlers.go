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

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
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

func convertUpdateUserResponse(user models.User) web.UpdateUserResponse {
	return web.UpdateUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: *user.UpdatedAt,
	}
}

func convertReturnUserResponse(ctx *gin.Context, code int, message string, data interface{}) {
	switch code != 0 {
	case code == 201:
		ctx.JSON(code, gin.H{
			"message": message,
			"user":    data,
		})
	case code == 200 && message == "Login Success":
		ctx.JSON(code, gin.H{
			"message": message,
			"token":   data,
		})
	case code == 200:
		ctx.JSON(code, gin.H{
			"message": message,
			"user":    data,
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

	userResp := &user

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
		helpers.ConvertErrResponse(ctx, http.StatusUnauthorized, "Password Failed", "Unauthenthicated")
		return
	}

	genToken := helpers.GenerateToken(user.ID, user.Email)

	convertReturnUserResponse(ctx, http.StatusOK, "Login Success", genToken)
}

func (h *UserHandlerImpl) GetUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	user, err := h.UserService.GetUser(userID)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Not Found", err.Error())
		return
	}

	convertReturnUserResponse(ctx, http.StatusOK, "User Found", user)
}

func (h *UserHandlerImpl) UpdateUser(ctx *gin.Context) {
	var updateInput web.UpdateUserRequest
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	_, err := h.UserService.GetUser(userID)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "User Not Found", err.Error())
	}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&updateInput)
	} else {
		ctx.ShouldBind(&updateInput)
	}

	updateUser, err := h.UserService.UpdateUser(updateInput, userID)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusBadRequest, "User Failed Update", err.Error())
		return
	}

	updateResp := &updateUser
	convertReturnUserResponse(ctx, http.StatusOK, "Updated User", updateResp)
}

func (h *UserHandlerImpl) DeleteUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	_, err := h.UserService.GetUser(userID)
	err = h.UserService.Delete(userID)

	if err != nil {
		helpers.ConvertErrResponse(ctx, http.StatusNotFound, "User Not Found", err.Error())
		return
	}

	convertReturnUserResponse(ctx, http.StatusAccepted, "User has ben deleted", "User has ben deleted")

}
