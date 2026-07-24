package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// instantiate a gin web server
	server := gin.Default()

	// routes
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// start the server running and listening
	server.Run(":8080") // http://localhost:8080
}

func getEvents(context *gin.Context) {
	allEvents := models.GetAllEvents()
	context.JSON(http.StatusOK, allEvents)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse event from request body"})
		return
	}

	// TODO: make this come from the db etc.
	event.ID = 1
	event.UserID = 1

	// save the event to persistence
	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
