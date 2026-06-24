package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/findamovieforme/users-service/models"
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

func GetMostAddedMovies(client *dynamodb.Client) ([]models.MostAddedMovie, error) {
	result, err := client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("MoviePopularity"),
		IndexName:              aws.String("PopularityIndex"),
		KeyConditionExpression: aws.String("#type = :type"),
		ExpressionAttributeNames: map[string]string{
			"#type": "type",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":type": &types.AttributeValueMemberS{Value: "POPULARITY"},
		},
		ScanIndexForward: aws.Bool(false),
		Limit:            aws.Int32(10),
	})
	if err != nil {
		return nil, err
	}

	var mostAddedMovies []models.MostAddedMovie
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &mostAddedMovies); err != nil {
		return nil, err
	}

	return mostAddedMovies, nil
}
