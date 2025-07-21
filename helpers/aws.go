package helpers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InitDynamoDBClient() *dynamodb.Client {
	region, err := LoadEnv("AWS_REGION")
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	return dynamodb.NewFromConfig(cfg)
}
