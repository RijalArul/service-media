package middlewares

import (
	"net/http"
	"service-media/databases"
	"service-media/helpers"
	models "service-media/models/entity"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		photo := models.Photo{}
		photoId := ctx.Param("photoId")
		photoParse, _ := strconv.ParseUint(photoId, 10, 32)
		err := db.Model(&photo).First(&photo, photoParse).Error

		if err != nil {
			helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Photo Not Found", err.Error())
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if photo.UserID != userID {
			helpers.ConvertErrResponse(ctx, http.StatusUnauthorized, "Unauthorized", "Unauthorized")
			return
		}

		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		comment := models.Comment{}
		commentId := ctx.Param("commentId")
		commentParse, _ := strconv.ParseUint(commentId, 10, 32)
		err := db.Model(&comment).First(&comment, commentParse).Error

		if err != nil {
			helpers.ConvertErrResponse(ctx, http.StatusNotFound, "Comment Not Found", err.Error())
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if comment.UserID != userID {
			helpers.ConvertErrResponse(ctx, http.StatusUnauthorized, "Unauthorized", "Unauthorized")
			return
		}

		ctx.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		socialMedia := models.SocialMedia{}
		socialMediaId := ctx.Param("socialMediaId")
		socialMediaParse, _ := strconv.ParseUint(socialMediaId, 10, 32)
		err := db.Model(&socialMedia).First(&socialMedia, socialMediaParse).Error

		if err != nil {
			helpers.ConvertErrResponse(ctx, http.StatusNotFound, "socialMedia Not Found", err.Error())
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if socialMedia.UserID != userID {
			helpers.ConvertErrResponse(ctx, http.StatusUnauthorized, "Unauthorized", "Unauthorized")
			return
		}

		ctx.Next()
	}
}
