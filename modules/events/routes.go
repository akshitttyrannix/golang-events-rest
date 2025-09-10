package events

import (
	"example.com/events/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.Use(middlewares.AuthMiddleware)

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
}
