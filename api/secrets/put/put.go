package main

import (
	"context"
	"fmt"
	"garage-vault/api/models"
	"garage-vault/api/utils"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Prepare the item to be added
	var body models.SecretEntry
	if err := utils.ValidateBody(request.Body, &body); err != nil {
		return nil, err
	}
	body.Pk = "pk_secret"

	// Put item in DynamoDB
	table := os.Getenv("DYNAMO_TABLE")

	db := dynamodb.New(session)
	item, err := dynamodbattribute.MarshalMap(&body)
	fmt.Printf("inserting item %v\n", item)
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
	utils.LambdaStart(HandleRequest)
}
