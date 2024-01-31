package main

import (
	"chat-server/config"
	"chat-server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to db
	config.Connect()

	// create a server with gin
	router := gin.Default()

	// middlewares, cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// routes
	routes.Routes(router)

	router.Run("localhost:9090") // path, port
}
