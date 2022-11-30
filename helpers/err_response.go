package helpers

import "github.com/gin-gonic/gin"

func ConvertErrResponse(ctx *gin.Context, code int, message string, err interface{}) {
	switch code != 0 {
	case code == 400:
		ctx.JSON(code, gin.H{
			"message": message,
			"err":     err,
		})
	case code == 401:
		ctx.JSON(code, gin.H{
			"message": "message",
			"err":     err,
		})
	case code == 404:
		ctx.JSON(code, gin.H{
			"message": "message",
			"err":     err,
		})
	default:
		ctx.JSON(code, gin.H{
			"message": "Internal Server Error",
		})
	}
}
