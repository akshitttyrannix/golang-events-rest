package users

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.POST("/users/sign-up", signUpUser)
	router.POST("/users/login", loginUser)
}
