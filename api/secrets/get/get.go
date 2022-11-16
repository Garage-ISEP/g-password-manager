package main

import (
	"context"
	"garage-vault/api/utils"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	table := os.Getenv("DYNAMO_TABLE")
	// Create DynamoDB db
	db := dynamodb.New(session)
	items, err := db.Query(&dynamodb.QueryInput{
		TableName: &table,
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		res := utils.HandleAWSProxy(HandleRequest, ctx, request)
		return res, nil
	})
}
