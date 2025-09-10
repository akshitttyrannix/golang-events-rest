package events

import (
	"errors"
	"time"

	"example.com/events/common/error"
	"example.com/events/common/messages"
	"example.com/events/common/success"
	"example.com/events/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createEvent(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		error.Unauthorized(ctx, errors.New("Unauthorized"))
		return
	}

	_, userID, err := utils.VerifyToken(token)
	if err != nil {
		error.Unauthorized(ctx, err)
		return
	}

	var event Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	if event.UserID != userID {
		error.Unauthorized(ctx, errors.New("Unauthorized"))
		return
	}

	event.EventID = uuid.New().String()
	event.CreatedAt = time.Now().Unix()
	event.UpdatedAt = time.Now().Unix()

	createdEvent, err := event.save()

	if err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_CREATED, createdEvent)
}

func getEvents(ctx *gin.Context) {
	allEvents, err := find()
	if err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENTS_FETCHED, allEvents)
}

func getEventByID(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := findByID(id)
	if err != nil {
		error.NotFound(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_FETCHED, event)
}

func updateEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := findByID(id)
	if err != nil {
		error.NotFound(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&event); err != nil {
		error.BadRequest(ctx, err)
		return
	}

	event.UpdatedAt = time.Now().Unix()

	updatedEvent, err := updateOne(event)
	if err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_UPDATED, updatedEvent)
}

func deleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := findByID(id)
	if err != nil {
		error.NotFound(ctx, err)
		return
	}

	err = deleteOne(id)
	if err != nil {
		error.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_DELETED, event)
}
