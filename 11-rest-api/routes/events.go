package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func HandleEventsGet(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func HandleEventCreate(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse event from request body"})
		return
	}

	// TODO: make this come from the db etc.
	event.UserID = 1

	// save the event to persistence
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func HandleEventGet(context *gin.Context) {

	// verify an Id was passed in
	eventIdAsString, paramFound := context.Params.Get("id")
	if !paramFound {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "id is a required field",
		})
		return
	}

	// verify the Id was numeric
	eventId, err := strconv.ParseInt(eventIdAsString, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(
			"Unable to convert id '%v' to numeric value.", eventIdAsString),
		})
		return
	}

	// perform the query operation
	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to get fetch event.",
		})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to find event from given EventId",
		})
		return
	}

	// happy path
	context.JSON(http.StatusOK, *event)
}

func HandleEventPut(context *gin.Context) {
	// verify an Id was passed in
	eventIdAsString, paramFound := context.Params.Get("id")
	if !paramFound {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "id is a required field",
		})
		return
	}

	// verify the Id was numeric
	eventId, err := strconv.ParseInt(eventIdAsString, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(
			"Unable to convert id {%v} to numeric value.", eventIdAsString),
		})
		return
	}

	// get the existingEvent that's being updated
	existingEvent, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to fetch event",
		})
		return
	}

	if existingEvent == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to find event",
		})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body",
		})
		return
	}

	// finally before the update
	newEvent, err := models.UpdateEvent(eventId, updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse request body",
		})
		return
	}

	// happy path
	context.JSON(http.StatusOK, newEvent)

}

func HandleEventDelete(context *gin.Context) {
	// verify an Id was passed in
	eventIdAsString, paramFound := context.Params.Get("id")
	if !paramFound {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "id is a required field",
		})
		return
	}

	// verify the Id was numeric
	eventId, err := strconv.ParseInt(eventIdAsString, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(
			"Unable to convert id '%v' to numeric value.", eventIdAsString),
		})
		return
	}

	// get the existingEvent that's being updated
	existingEvent, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to fetch event",
		})
		return
	}

	if existingEvent == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to find event",
		})
		return
	}

	err = models.DeleteEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete event",
		})
	}

	// happy path
	context.JSON(http.StatusNoContent, gin.H{
		"message": "event successfully deleted",
	})

}
