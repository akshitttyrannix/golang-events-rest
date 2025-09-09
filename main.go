package main

import (
	"example.com/events/common/database"
	"example.com/events/modules/events"
	"example.com/events/modules/users"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	server := gin.Default()

	prefix := server.Group("/api/v1")
	events.Routes(prefix)
	users.Routes(prefix)

	server.Run(":8080")
}
