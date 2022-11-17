package main

import (
	"context"
	"garage-vault/api/models"
	"garage-vault/api/utils"
	"os"
	"time"

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
	body.Pk = models.PK_SECRET
	body.Sk = body.Group + "#" + body.Name
	// body.CreatedById = request.RequestContext.Authorizer["principalId"].(string)
	// body.CreatedByName = request.RequestContext.Authorizer["name"].(string)
	body.CreatedAt = time.Now().Format(time.RFC3339)
	body.UpdatedAt = body.CreatedAt
	// body.UpdatedById = body.CreatedById
	// body.UpdatedByName = body.CreatedByName

	// Put item in DynamoDB
	table := os.Getenv("DYNAMO_TABLE")

	db := dynamodb.New(session)
	item, err := dynamodbattribute.MarshalMap(&body)
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

	var output models.SecretEntry
	if err := dynamodbattribute.UnmarshalMap(item, &output); err != nil {
		return nil, err
	}
	return output, nil
}

func main() {
	utils.LambdaStart(HandleRequest)
}
