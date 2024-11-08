package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	usersGroup := router.Group("/users")
	usersGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Users service started successfully!",
		})
	})
	usersGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router
}
