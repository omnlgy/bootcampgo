package router

import (
	"time"

	"example.com/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisteredRoute(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(utils.ApiAuthMiddleware())
	api.Use(utils.ApiLimiterMiddleware(utils.ApiLimiterConfig{
		MaxRequests: 10,
		Window:      1 * time.Minute,
	}))

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})
}
