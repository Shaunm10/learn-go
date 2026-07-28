package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// event routes
	server.GET("/events", HandleEventsGet)
	server.GET("/events/:id", HandleEventGet)

	// these should be protected by authentication
	server.POST("/events", HandleEventCreate)
	server.PUT("/events/:id", HandleEventPut)
	server.DELETE("/events/:id", HandleEventDelete)

	// user routes
	server.POST("/signup", HandleSignupPOST)
	server.POST("/login", HandleLoginPOST)
}
