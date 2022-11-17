package main

import (
	"context"
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
	Query string `json:"query"`
	From  string `json:"from"`
}

type Response struct {
	Items *[]models.SecretEntry `json:"items"`
	Query string                `json:"query"`
	From  string                `json:"from"`
	To    string                `json:"to"`
	Total int64                 `json:"total"`
	Count int64                 `json:"count"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queryParams := Request{
		Query: request.QueryStringParameters["query"],
		From:  request.QueryStringParameters["from"],
	}
	if queryParams.Query == "" {
		queryParams.Query = ".*"
	}

	table := os.Getenv("DYNAMO_TABLE")
	// Create DynamoDB db
	db := dynamodb.New(session)

	queryRes, err := db.Query(&dynamodb.QueryInput{
		TableName:              &table,
		Limit:                  aws.Int64(50),
		KeyConditionExpression: aws.String("contains(sk, :query"),
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"sk": {
				S: aws.String(queryParams.From),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":query": {
				S: aws.String(queryParams.Query),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var items *[]models.SecretEntry
	if err := dynamodbattribute.UnmarshalListOfMaps(queryRes.Items, items); err != nil {
		return nil, err
	}

	return &Response{
		Items: items,
		Count: *queryRes.Count,
		Total: *queryRes.ScannedCount,
		From:  queryParams.From,
		To:    *queryRes.LastEvaluatedKey["sk"].S,
	}, nil
}

func main() {
	utils.LambdaStart(HandleRequest)
}
