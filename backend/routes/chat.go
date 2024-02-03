package routes

import (
	"chat-server/controller"

	"github.com/gin-gonic/gin"
)

func ChatRouter(router *gin.Engine) {
	g := router.Group("/api")
	g.GET("/chat", controller.GetChats)
	g.GET("/ws", controller.WebSocketConnect)
}
