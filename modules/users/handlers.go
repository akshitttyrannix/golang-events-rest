package users

import (
	"time"

	"example.com/events/common/errors"
	"example.com/events/common/messages"
	"example.com/events/common/success"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func signUpUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		errors.BadRequest(ctx, err)
		return
	}

	user.UserID = uuid.New().String()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	createdUser, err := user.save()
	if err != nil {
		errors.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.USER_SIGNED_UP, createdUser)
}

func loginUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		errors.BadRequest(ctx, err)
		return
	}

	if err := checkCredentials(user.Email, user.Password); err != nil {
		errors.Unauthorized(ctx, err)
		return
	}

	success.Success(ctx, messages.USER_LOGGED_IN, nil)
}
