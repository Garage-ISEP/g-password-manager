package main

import (
	"context"
	"fmt"
	"garage-vault/api/models"
	"garage-vault/api/utils"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Request struct {
	Query string
	From  string
}

type Response struct {
	Items *[]models.SecretEntry `json:"items"`
	Query *string               `json:"query"`
	From  *string               `json:"from"`
	To    *string               `json:"to"`
	Total *int64                `json:"total"`
	Count *int64                `json:"count"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var queryParams Request
	if err := utils.ValidateQueryParams(&request.QueryStringParameters, &queryParams); err != nil {
		return nil, err
	}

	table := os.Getenv("DYNAMO_TABLE")
	// Create DynamoDB db
	db := dynamodb.New(session)

	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(table),
		Limit:                  aws.Int64(50),
		KeyConditionExpression: aws.String("pk = :pk"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String(models.PK_SECRET),
			},
		},
	}
	// Handle from query offset
	if queryParams.From != "" {
		queryInput.SetExclusiveStartKey(map[string]*dynamodb.AttributeValue{
			"sk": {
				S: aws.String(queryParams.From),
			},
		})
	}
	// Handler query string to search for a specific secret
	if queryParams.Query != "" {
		queryInput.SetFilterExpression("contains(#name, :query)")
		queryInput.ExpressionAttributeValues[":query"] = &dynamodb.AttributeValue{
			S: aws.String(queryParams.Query),
		}
	}
	// Execute query
	queryRes, err := db.Query(queryInput)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}

	// Parse response
	var items []models.SecretEntry = make([]models.SecretEntry, 0)
	if err := dynamodbattribute.UnmarshalListOfMaps(queryRes.Items, &items); err != nil {
		return nil, err
	}

	var to *string
	if queryRes.LastEvaluatedKey != nil {
		to = queryRes.LastEvaluatedKey["sk"].S
	}
	return &Response{
		Items: &items,
		Count: queryRes.Count,
		Total: queryRes.ScannedCount,
		From:  &queryParams.From,
		To:    to,
	}, nil
}

func main() {
	utils.LambdaStart(HandleRequest)
}
