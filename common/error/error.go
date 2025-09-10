package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func NotFound(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func SomethingWentWrong(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func Unauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func UnprocessableEntity(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func Abort(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}
