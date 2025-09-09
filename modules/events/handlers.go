package events

import (
	"time"

	"example.com/events/common/errors"
	"example.com/events/common/messages"
	"example.com/events/common/success"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createEvent(ctx *gin.Context) {
	var event Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		errors.BadRequest(ctx, err)
		return
	}

	event.EventID = uuid.New().String()
	event.CreatedAt = time.Now().Unix()
	event.UpdatedAt = time.Now().Unix()

	createdEvent, err := event.save()

	if err != nil {
		errors.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_CREATED, createdEvent)
}

func getEvents(ctx *gin.Context) {
	allEvents, err := find()
	if err != nil {
		errors.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENTS_FETCHED, allEvents)
}

func getEventByID(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := findByID(id)
	if err != nil {
		errors.NotFound(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_FETCHED, event)
}

func updateEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := findByID(id)
	if err != nil {
		errors.NotFound(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&event); err != nil {
		errors.BadRequest(ctx, err)
		return
	}

	event.UpdatedAt = time.Now().Unix()

	updatedEvent, err := updateOne(event)
	if err != nil {
		errors.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_UPDATED, updatedEvent)
}

func deleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := findByID(id)
	if err != nil {
		errors.NotFound(ctx, err)
		return
	}

	err = deleteOne(id)
	if err != nil {
		errors.SomethingWentWrong(ctx, err)
		return
	}

	success.Success(ctx, messages.EVENT_DELETED, event)
}
