package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) ([]*string, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(session)
	tables, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	return tables.TableNames, nil
}

func main() {
	lambda.Start(HandleRequest)
}
