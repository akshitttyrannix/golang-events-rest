package events

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
}
