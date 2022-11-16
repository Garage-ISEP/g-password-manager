package main

import (
	"context"
	"garage-vault/api/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
	return nil, nil
}

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		res := utils.HandleAWSProxy(HandleRequest, ctx, request)
		return res, nil
	})
}
