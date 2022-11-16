package main

import (
	"context"
	"encoding/json"
	"garage-vault/api/utils"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type RequestBody struct {
	Name        string `json:"name"`
	Secret      string `json:"secret"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Group       string `json:"group"`
	Description string `json:"description"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var body RequestBody
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return nil, err
	}

	table := os.Getenv("DYNAMO_TABLE")

	db := dynamodb.New(session)
	item, err := dynamodbattribute.MarshalMap(&request)
	if err != nil {
		return nil, err
	}

	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName: &table,
		Item:      item,
	})
	if err != nil {
		return nil, err
	}

	return item, nil
}

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		res := utils.HandleAWSProxy(HandleRequest, ctx, request)
		return res, nil
	})
}
