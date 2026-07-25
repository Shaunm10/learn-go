package main

import (
	"net/http"

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
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// start the server running and listening
	server.Run(":8080") // http://localhost:8080

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
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
