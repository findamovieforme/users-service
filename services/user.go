package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/movierecuh/users-service/models"
)

func StoreUserPreferences(client *dynamodb.Client, userID string, movies interface{}) error {
	// Marshal preferences with interface{}
	preferences, err := attributevalue.Marshal(movies)
	if err != nil {
		return err
	}

	// Prepare the item for DynamoDB
	item := map[string]types.AttributeValue{
		"user_id":     &types.AttributeValueMemberS{Value: userID},
		"preferences": preferences,
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("UserPreferences"),
		Item:      item,
	})

	return err
}

func GetUserPreferences(client *dynamodb.Client, userID string) (*models.UserPreferences, error) {
	result, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("UserPreferences"),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: userID},
		},
	})

	if err != nil {
		return nil, err
	}

	// Check if item exists; if not, return ErrNotFound
	if result.Item == nil {
		return nil, models.ErrNotFound
	}
	// Unmarshal into UserPreferences
	var userPreferences models.UserPreferences
	err = attributevalue.UnmarshalMap(result.Item, &userPreferences)
	if err != nil {
		return nil, err
	}

	return &userPreferences, nil
}

func GetMostAddedMovies(client *dynamodb.Client) ([]interface{}, error) {
	// Aggregate all the users to get the most added movies
	result, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("UserPreferences"),
	})

	if err != nil {
		return nil, err
	}

	// Count the number of times each movie is added
	movieCount := make(map[float64]int)
	moiveDetailsMap := make(map[float64]interface{})
	for _, item := range result.Items {
		var userPreferences models.UserPreferences
		err = attributevalue.UnmarshalMap(item, &userPreferences)
		if err != nil {
			return nil, err
		}

		for _, movie := range userPreferences.Preferences {
			movieMap := movie.(map[string]interface{})
			movieTitle := movieMap["id"].(float64)
			movieCount[movieTitle]++
			moiveDetailsMap[movieTitle] = movieMap
		}
	}

	// Find the most added movies
	var mostAddedMovies []interface{}
	for movie, count := range movieCount {
		mostAddedMovies = append(mostAddedMovies, map[string]interface{}{
			"id":    movie,
			"count": count,
			// Return the entire movie object
			"details": moiveDetailsMap[movie],
		})
	}

	return mostAddedMovies, nil
}
