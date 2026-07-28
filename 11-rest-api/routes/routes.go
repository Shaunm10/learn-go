package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// event routes
	server.GET("/events", HandleEventsGet)
	server.GET("/events/:id", HandleEventGet)

	// these should be protected by authentication
	authenticatedRouterGroup := server.Group("/")

	// adds this middlewear to all the routes mentioned below
	authenticatedRouterGroup.Use(middleware.Authenticate)
	authenticatedRouterGroup.POST("/events", HandleEventCreate)
	authenticatedRouterGroup.PUT("/events/:id", HandleEventPut)
	authenticatedRouterGroup.DELETE("/events/:id", HandleEventDelete)

	// user routes
	server.POST("/signup", HandleSignupPOST)
	server.POST("/login", HandleLoginPOST)
}
