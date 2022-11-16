package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	return fmt.Sprintf("Hello PUT %s!", request.Body), nil
}

func main() {
	lambda.Start(HandleRequest)
}
