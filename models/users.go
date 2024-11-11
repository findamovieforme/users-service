package models

import "errors"

type UserPreferences struct {
	UserID      string        `json:"user_id"`
	Preferences []interface{} `json:"preferences"` // Use interface{} for flexible JSON data
}

// Define a custom error for not found cases
var ErrNotFound = errors.New("item not found")
