package routes

import "github.com/gin-gonic/gin"

func PingRoute(r *gin.Engine) {
	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
