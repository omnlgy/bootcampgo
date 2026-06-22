package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

func ApiAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("api-key")
		if token != "supersecretkey" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

type ApiLimiterConfig struct {
	MaxRequests int
	Window      time.Duration
}

var userRequests = make(map[string]struct {
	hitCount  int
	lastReset time.Time
})

func ApiLimiterMiddleware(config ApiLimiterConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIp := c.ClientIP()
		if _, exists := userRequests[userIp]; !exists {
			userRequests[userIp] = struct {
				hitCount  int
				lastReset time.Time
			}{
				hitCount:  0,
				lastReset: time.Now(),
			}
		} else {
			userRequests[userIp] = struct {
				hitCount  int
				lastReset time.Time
			}{
				hitCount:  userRequests[userIp].hitCount + 1,
				lastReset: userRequests[userIp].lastReset,
			}

			if time.Since(userRequests[userIp].lastReset) > config.Window {
				delete(userRequests, userIp)
			}
		}

		if userRequests[userIp].hitCount > config.MaxRequests {
			c.JSON(429, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
