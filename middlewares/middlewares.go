package middlewares

import (
	"errors"

	"example.com/events/common/error"
	"example.com/events/common/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		error.Abort(ctx, errors.New("Unauthorized"))
		return
	}

	_, userID, err := utils.VerifyToken(token)
	if err != nil {
		error.Abort(ctx, err)
		return
	}

	ctx.Set("user_id", userID)

	ctx.Next()
}
