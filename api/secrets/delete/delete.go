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
)

type Request struct {
	Sk *string `required:"true"`
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

	db := dynamodb.New(session)

	return db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(models.PK_SECRET),
			},
			"sk": {
				S: queryParams.Sk,
			},
		},
		ConditionExpression: aws.String("contains(sk, :sk)"),
	})
}

func main() {
	utils.LambdaStart(HandleRequest)
}
