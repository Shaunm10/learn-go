package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", HandleEventsGet)
	server.POST("/events", HandleEventCreate)
	server.GET("/events/:id", HandleEventGet)
	server.PUT("/events/:id", HandleEventPut)
	server.DELETE("/events/:id", HandleEventDelete)
}
