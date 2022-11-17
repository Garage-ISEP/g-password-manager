package main

import (
	"context"
	"garage-vault/api/utils"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
	return nil, nil
}

func main() {
	utils.LambdaStart(HandleRequest)
}
