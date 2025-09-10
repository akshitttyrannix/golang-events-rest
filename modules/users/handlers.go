package users

import (
	"time"

	"example.com/events/common/error"
	"example.com/events/common/messages"
	"example.com/events/common/success"
	"example.com/events/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func signUpUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	user.UserID = uuid.New().String()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	createdUser, err := user.save()
	if err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.USER_SIGNED_UP, createdUser)
}

func loginUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	if err := user.checkCredentials(); err != nil {
		error.Unauthorized(ctx, err)
		return
	}

	token, err := utils.GenerateToken(user.UserID, user.Email)
	if err != nil {
		error.UnprocessableEntity(ctx, err)
		return
	}

	success.Success(ctx, messages.USER_LOGGED_IN, token)
}
