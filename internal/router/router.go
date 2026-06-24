package router

import (
	"example.com/internal/controller"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(server *gin.Engine, controller *controller.OrderController) {
	api := server.Group("/api")

	api.POST("/orders", controller.CreateOrder)

}
