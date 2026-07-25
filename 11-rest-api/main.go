package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/database"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	// start the db connection
	database.InitDB()

	// instantiate a gin web server
	server := gin.Default()

	// routes
	server.GET("/events", handleGetEventsRequest)
	server.POST("/events", handleCreateEventRequest)
	server.GET("/events/:id", handleGetEventRequest)

	// start the server running and listening
	server.Run(":8080") // http://localhost:8080

}

func handleGetEventsRequest(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func handleCreateEventRequest(context *gin.Context) {
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

func handleGetEventRequest(context *gin.Context) {

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
