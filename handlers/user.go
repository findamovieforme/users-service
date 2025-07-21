package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/findamovieforme/users-service/models"
	"github.com/findamovieforme/users-service/services"
)

type UserHandler struct {
	dynamoDBClient *dynamodb.Client
}

func NewUserHandler(dynamoDBClient *dynamodb.Client) *UserHandler {
	return &UserHandler{
		dynamoDBClient: dynamoDBClient,
	}
}

// SaveUserPreferences - Gin handler to save user preferences
func (h *UserHandler) SaveUserPreferences(c *gin.Context) {
	// Bind JSON request body to a struct
	var req struct {
		UserID string        `json:"user_id"`
		Movies []interface{} `json:"movies"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call a function to save preferences to DynamoDB
	if err := services.StoreUserPreferences(h.dynamoDBClient, req.UserID, req.Movies); err != nil {
		log.Println("Failed to save preferences:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save preferences"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Preferences saved successfully"})
}

func (h *UserHandler) FetchUserPreference(c *gin.Context) {
	userID := c.Request.URL.Query().Get("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Fetch user preferences
	userPreferences, err := services.GetUserPreferences(h.dynamoDBClient, userID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			// Handle not found case with 404 response
			c.JSON(http.StatusNotFound, gin.H{"error": "User preferences not found"})
			return
		}
		// Handle other errors
		fmt.Println("Error fetching user preferences:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Return the user preferences
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "preferences": userPreferences.Preferences})
}

func (h *UserHandler) FetchMostAdded(c *gin.Context) {
	// Fetch the most added movies
	mostAddedMovies, err := services.GetMostAddedMovies(h.dynamoDBClient)
	if err != nil {
		fmt.Println("Error fetching most added movies:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": mostAddedMovies})
}
