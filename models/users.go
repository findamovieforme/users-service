package models

import "errors"

type UserPreferences struct {
	UserID      string        `json:"user_id"`
	Preferences []interface{} `json:"preferences"` // Use interface{} for flexible JSON data
}

type MostAddedMovie struct {
	ID    int    `json:"id" dynamodbav:"movie_id"`
	Count int    `json:"count" dynamodbav:"count"`
	Type  string `json:"type" dynamodbav:"type"`
}

// Define a custom error for not found cases
var ErrNotFound = errors.New("item not found")
