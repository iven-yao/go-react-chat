package routes

import (
	"chat-server/controller"

	"github.com/gin-gonic/gin"
)

func ChatRouter(router *gin.Engine) {
	g := router.Group("/api")
	g.GET("/chat/:chatroomId", controller.GetChats)
	g.POST("/chat", controller.SendChat)
}
