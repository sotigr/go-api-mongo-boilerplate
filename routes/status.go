package routes

import (
	controller "main/controllers/default"

	"github.com/gin-gonic/gin"
)

func StatusRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("status", controller.GetStatus())
}
