package router

import (
	"example.com/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, orderCtrl *controller.OrderController, productCtrl *controller.ProductController) {
	api := server.Group("/api")

	api.GET("/orders", orderCtrl.GetOrders)
	api.POST("/orders", orderCtrl.CreateOrder)

	api.POST("/products", productCtrl.CreateProduct)
}

func RegisterAuthRoutes(server *gin.Engine, authCtrl *controller.AuthController) {
	api := server.Group("/api")

	api.POST("/login", authCtrl.LoginHandler)
	api.POST("/register", authCtrl.RegisterHandler)
}
