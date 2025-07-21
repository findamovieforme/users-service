package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/findamovieforme/users-service/handlers"
	"github.com/findamovieforme/users-service/helpers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

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

	dynamoDBClient := helpers.InitDynamoDBClient()
	userHandler := handlers.NewUserHandler(dynamoDBClient)

	// Define routes and their corresponding handlers
	usersGroup.GET("/preferences", userHandler.FetchUserPreference)
	usersGroup.POST("/preferences", userHandler.SaveUserPreferences)

	usersGroup.GET("/mostAdded", userHandler.FetchMostAdded)

	return router
}
