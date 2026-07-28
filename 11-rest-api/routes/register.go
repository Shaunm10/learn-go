package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func HandleRegisterCreate(context *gin.Context) {
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

	authenticatedUserId := context.GetInt64("userId")

	// get the event so we can verify it exists
	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to get event",
		})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to find event",
		})
	}

	register, err := models.CreateRegister(authenticatedUserId, eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create register"})
		return
	}

	context.JSON(http.StatusCreated, register)
}

func HandleRegisterDelete(context *gin.Context) {
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

	authenticatedUserId := context.GetInt64("userId")

	err = models.DeleteRegister(authenticatedUserId, eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error removing registration",
		})
		return
	}

	// happy path
	context.JSON(http.StatusNoContent, gin.H{
		"message": "event successfully deleted",
	})

}
