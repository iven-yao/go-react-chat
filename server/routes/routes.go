package routes

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	UserRoutes(router)
	ChatRouter(router)
}
