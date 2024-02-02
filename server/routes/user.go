package routes

import (
	"chat-server/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	g := router.Group("/api")
	g.GET("/user", controller.GetUsers)
	g.POST("/user/register", controller.Register)
	g.POST("/user/login", controller.Login)
	g.GET("/user/test", controller.TestUser)
}
