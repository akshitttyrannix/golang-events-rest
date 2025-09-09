package success

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
		"status":  1,
	})
}
