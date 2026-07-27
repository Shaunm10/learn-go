package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// event routes
	server.GET("/events", HandleEventsGet)
	server.POST("/events", HandleEventCreate)
	server.GET("/events/:id", HandleEventGet)
	server.PUT("/events/:id", HandleEventPut)
	server.DELETE("/events/:id", HandleEventDelete)

	// user routes
	server.POST("/signup", HandleSignupPOST)
	server.POST("/login", HandleLoginPOST)
}
