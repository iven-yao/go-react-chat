package routes

import (
	"chat-server/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	g := router.Group("/api/user")
	g.GET("", controller.GetUsers)
	g.POST("", controller.CreateUser)
	g.GET("/:id", controller.GetUserById)
	g.DELETE("/:id", controller.DeleteUser)
	g.PUT("/:id", controller.UpdateUser)
}
