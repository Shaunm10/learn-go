package main

import (
	"example.com/rest-api/database"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// start the db connection
	database.InitDB()

	// instantiate a gin web server
	server := gin.Default()

	// routes
	routes.RegisterRoutes(server)

	// start the server running and listening
	server.Run(":8080") // http://localhost:8080
}
